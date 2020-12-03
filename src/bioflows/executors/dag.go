package executors

import (
	config2 "bioflows/config"
	"bioflows/managers"
	"bioflows/models"
	"bioflows/models/pipelines"
	"bioflows/resolver"
	"errors"
	"fmt"
	"github.com/goombaio/dag"
	"log"
	"os"
	"strings"
	"sync"
)

type DagExecutor struct {
	contextManager *managers.ContextManager
	planManager *managers.ExecutionPlanManager
	transformations []TransformCall
	parentPipeline *pipelines.BioPipeline
	logger *log.Logger
	containerConfig *models.ContainerConfig
	scheduler *DagScheduler
	rankedList [][]*dag.Vertex
}
func (p *DagExecutor) SetContainerConfig(containerConfig *models.ContainerConfig) {
	p.containerConfig = containerConfig
}
func (p *DagExecutor) SetContext(c *managers.ContextManager) {
	p.contextManager = c
}

func (p *DagExecutor) copyParentParamsInto(step *pipelines.BioPipeline) {
	if len(p.parentPipeline.Inputs) > 0 {
		if step.Inputs == nil || len(step.Inputs) == 0{
			step.Inputs = make([]models.Parameter,len(p.parentPipeline.Inputs))
			copy(step.Inputs,p.parentPipeline.Inputs)
		}else{
			step.Inputs = append(step.Inputs,p.parentPipeline.Inputs...)
		}
	}
}
func (p *DagExecutor) GetContext() *managers.ContextManager {
	return p.contextManager
}
//This function returns the final result of the current pipeline
func (p *DagExecutor) GetPipelineOutput() models.FlowConfig {
	tempConfig := models.FlowConfig{}
	pipelineKey := resolver.ResolvePipelineKey(p.parentPipeline.ID)
	pipelineConfig , err := p.GetContext().GetStateManager().GetPipelineState(pipelineKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to fetch Pipeline Configuration for %s",pipelineKey))
		return tempConfig
	}
	tempConfig.Fill(pipelineConfig)
	return tempConfig
}


func (p DagExecutor) SetPipelineGeneralConfig(b *pipelines.BioPipeline,originalConfig *models.FlowConfig) {
	// Read the pipeline general configuration section
	if b.Config != nil && len(b.Config) > 0 {
		internalConfig := make(map[string]interface{})
		for _ , param := range b.Config {
			internalConfig[param.Name] = param.Value
		}
		(*originalConfig)[config2.BIOFLOWS_INTERNAL_CONFIG] = internalConfig
	}
	//Attach the general container configuration if exists.
	if b.ContainerConfig != nil {
		p.containerConfig = b.ContainerConfig
	}
}
func (p *DagExecutor) Clean() bool {
	return p.contextManager.GetStateManager().RemoveConfigByID(resolver.BIOFLOWS_NAME)
}

func (p *DagExecutor) CheckStatus(pipelineId string , step pipelines.BioPipeline) int {
	status := SHOULD_RUN
	toolKey := resolver.ResolveToolKey(step.ID,pipelineId)
	section , _ := p.contextManager.GetStateManager().GetStateByID(toolKey)
	if section != nil {
		data := section.(map[string]interface{})
		if _, found := data["status"]; found {
			status = DONT_RUN
		}
	}
	//Check that all dependent steps have run successfully
	if len(step.Depends) > 0 {
		depends := strings.Split(step.Depends,",")
		result := true
		for _ , v := range depends {
			toolName := resolver.ResolveToolKey(v,pipelineId)
			data , _ := p.GetContext().GetStateManager().GetStateByID(toolName)
			if data != nil {
				toolConfig := data.(map[string]interface{})
				if statusVar , found := toolConfig["status"]; !found {
					status = SHOULD_QUEUE
				}else{
					result = result && (statusVar.(bool))
				}
			}else{
				status = SHOULD_QUEUE
			}

		}
		if !result{
			status = DONT_RUN
		}
	}
	return status
}

func (p *DagExecutor) Setup(config models.FlowConfig) error {

	p.scheduler = &DagScheduler{}

	p.transformations = make([]TransformCall,0)
	p.contextManager = &managers.ContextManager{}
	p.planManager = &managers.ExecutionPlanManager{}
	err := p.contextManager.Setup(config)
	if err != nil {
		return err
	}
	p.planManager.SetContextManager(p.contextManager)
	p.createLogFile(config)
	return p.planManager.Setup(config)
}
func (p *DagExecutor) createLogFile(config models.FlowConfig) error {
	workflowOutputFile := strings.Join([]string{
		fmt.Sprintf("%v",config[config2.WF_INSTANCE_OUTDIR]),
		"workflow.logs",
	},"/")
	p.logger = &log.Logger{}
	p.logger.SetPrefix(fmt.Sprintf("%v: ",config2.BIOFLOWS_NAME))
	file,  err := os.Create(workflowOutputFile)
	if err != nil {
		return err
	}
	p.logger.SetOutput(file)
	return nil
}

func (p *DagExecutor) Log(logs ...interface{}) {
	p.logger.Println(logs...)
	fmt.Println(logs...)
}

func (p *DagExecutor) Run(b *pipelines.BioPipeline,config models.FlowConfig) error {
	p.SetPipelineGeneralConfig(b,&config)
	var finalError error
	defer func() error{
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				finalError = r.(error)
				fmt.Println(fmt.Sprintf("Error: %s. Aborting.....",finalError.Error()))
			case string:
				finalError = errors.New(r.(string))
			default:
				finalError = errors.New("There was an exception while running the current pipeline....")
			}

		}
		return finalError
	}()
	p.parentPipeline = b
	finalError = p.runLocal(b,config)
	p.Log(fmt.Sprintf("Workflow: (%s) has finished....",b.Name))
	return finalError
}
func (p *DagExecutor) runLocal(b *pipelines.BioPipeline, config models.FlowConfig) error {
	graph , err := pipelines.CreateGraph(b)
	if err != nil {
		return err
	}
	p.rankedList , err = p.scheduler.Rank(b,graph)
	if err != nil {
		return err
	}
	if p.rankedList == nil {
		return errors.New("Failed to rank the current pipeline. Aborting....")
	}
	for _ , sublist := range p.rankedList {
		wg := sync.WaitGroup{}
		for _ , node := range sublist {
			if node == nil {
				continue
			}
			wg.Add(1)
			go p.execute(config,node,&wg)
		}
		wg.Wait()
	}
	return nil
}
func (p *DagExecutor) prepareConfig(b *pipelines.BioPipeline,config models.FlowConfig) models.FlowConfig {
	tempConfig := models.FlowConfig{}
	for k , v := range config{
		tempConfig[k] = v
	}
	pipelineKey := resolver.ResolvePipelineKey(b.ID)
	pipelineConfig , err := p.GetContext().GetStateManager().GetPipelineState(pipelineKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to fetch Pipeline Configuration for %s",pipelineKey))
		return tempConfig
	}
	tempConfig.Fill(pipelineConfig)
	return tempConfig
}
func (p *DagExecutor) execute(config models.FlowConfig,vertex *dag.Vertex,wg *sync.WaitGroup) {
	defer wg.Done()
	currentFlow := vertex.Value.(pipelines.BioPipeline)
	PreprocessPipeline(&currentFlow,config,p.transformations...)
	toolKey := resolver.ResolveToolKey(currentFlow.ID,p.parentPipeline.ID)
	//pipelineKey := resolver.ResolvePipelineKey(p.parentPipeline.ID)
	status := p.CheckStatus(p.parentPipeline.ID,currentFlow)
	switch status {
	case SHOULD_RUN:
		p.copyParentParamsInto(&currentFlow)
		if currentFlow.IsTool() {
			// It is a tool
			executor := ToolExecutor{}
			executor.SetPipelineName(p.parentPipeline.Name)
			executor.SetContainerConfiguration(p.containerConfig)
			toolInstance := &models.ToolInstance{
				WorkflowID: p.parentPipeline.ID,
				WorkflowName: p.parentPipeline.Name,
				Tool:currentFlow.ToTool(),
			}
			toolInstance.Prepare()
			generalConfig := p.prepareConfig(p.parentPipeline,config)
			toolInstanceFlowConfig , err := executor.Run(toolInstance,generalConfig)
			if err != nil {

				executor.Log(fmt.Sprintf("Received Error : %s",err.Error()))
			}
			if toolInstanceFlowConfig != nil {

				err = p.contextManager.SaveState(toolKey,toolInstanceFlowConfig.GetAsMap())
				if err != nil {
					fmt.Println(fmt.Sprintf("Received Error: %s",err.Error()))
					return
				}
			}
		}else{
			//it is a nested pipeline
			nestedPipelineExecutor := DagExecutor{}
			nestedPipelineExecutor.SetContainerConfig(p.containerConfig)
			nestedPipelineConfig := models.FlowConfig{}
			pipelineConfig := p.prepareConfig(&currentFlow,config)
			nestedPipelineConfig.Fill(config)
			nestedPipelineConfig.Fill(pipelineConfig)
			nestedPipelineExecutor.Setup(nestedPipelineConfig)
			err := nestedPipelineExecutor.Run(&currentFlow,nestedPipelineConfig)
			if err != nil {

				nestedPipelineExecutor.Log(err.Error())
			}
			pipeConfig := nestedPipelineExecutor.GetPipelineOutput()
			err = p.contextManager.SaveState(toolKey,pipeConfig)
		}
	case SHOULD_QUEUE:
		fallthrough
	case DONT_RUN:
		fallthrough
	default:
		p.Log(fmt.Sprintf("Flow: %s has already run before, deferring....",currentFlow.Name))
		return
	}
}




