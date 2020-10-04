package executors

import (
	config2 "bioflows/config"
	"bioflows/managers"
	"bioflows/models"
	"bioflows/models/pipelines"
	"bioflows/resolver"
	"fmt"
	"github.com/goombaio/dag"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type PipelineExecutor struct {
	contextManager *managers.ContextManager
	planManager *managers.ExecutionPlanManager
	transformations []TransformCall
	waitGroup sync.WaitGroup
	mutex *sync.Mutex
	waitQueue chan *dag.Vertex
	stopChan chan interface{}
	parentPipeline *pipelines.BioPipeline
	ticker *time.Ticker
	logger *log.Logger

}

//This function returns the final result of the current pipeline
func (p *PipelineExecutor) GetPipelineOutput() models.FlowConfig {
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

func (p *PipelineExecutor) IsRemote() bool {
	return p.contextManager.IsRemote()
}

func (p *PipelineExecutor) handleWaitQueue(config models.FlowConfig) {

	go func(){

		for{
			select{
			case <- p.ticker.C:
				if task , ok := <- p.waitQueue; ok {
					p.waitGroup.Add(1)
					p.executeSingleVertex(p.parentPipeline,config,task)

				}
			}
		}
	}()

}

func (p *PipelineExecutor) SetContext(c *managers.ContextManager) {
	p.contextManager = c
}

func (p *PipelineExecutor) GetContext() *managers.ContextManager {
	return p.contextManager
}
func (p *PipelineExecutor) Run(b *pipelines.BioPipeline,config models.FlowConfig) error {
	p.parentPipeline = b
	//Start handling wait queue
	p.handleWaitQueue(config)
	// Start processing the current pipeline
	PreprocessPipeline(b,config,p.transformations...)
	var finalError error
	if p.IsRemote(){
		finalError = p.runOnCluster(b,config)
	}else{
		finalError = p.runLocally(b,config)
	}
	p.Log(fmt.Sprintf("Workflow: (%s) has finished....",b.Name))
	return finalError

}
func (p *PipelineExecutor) canRun(pipelineId string , step pipelines.BioPipeline) bool {
	if len(step.Depends) <= 0 {
		return true
	}
	depends := strings.Split(step.Depends,",")
	result := true
	for _ , v := range depends {
		toolName := resolver.ResolveToolKey(v,pipelineId)
		_ , err := p.GetContext().GetStateManager().GetStateByID(toolName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s",err.Error()))
			result = false
			return result
		}
		//toolConfig := data.(map[string]interface{})
		//if status , ok := toolConfig["status"]; !ok {
		//	result = false
		//}else{
		//	result = result && !(status.(bool))
		//}
	}
	return result
}
func (p *PipelineExecutor) isAlreadyRun(toolKey string) bool{
	result := false
	section , err := p.contextManager.GetStateManager().GetStateByID(toolKey)
	if err != nil {
		result = false
		return result
	}
	if section == nil {
		return false
	}
	data := section.(map[string]interface{})
	if _ , ok := data["status"] ; ok {
		result = true
	}
	return result

}
func (p *PipelineExecutor) executeSingleVertex(b *pipelines.BioPipeline , config models.FlowConfig,vertex *dag.Vertex) {
	defer p.waitGroup.Done()
	currentFlow := vertex.Value.(pipelines.BioPipeline)
	toolKey := resolver.ResolveToolKey(currentFlow.ID,b.ID)
	//pipelineKey := resolver.ResolvePipelineKey(p.parentPipeline.ID)

	if p.canRun(b.ID,currentFlow) {
		if p.isAlreadyRun(toolKey){

			goto RunChildren
		}
		if currentFlow.IsTool() {
			// It is a single tool
			executor := ToolExecutor{}
			toolInstance := &models.ToolInstance{
				WorkflowID: b.ID,
				WorkflowName: b.Name,
				Tool:currentFlow.ToTool(),
			}
			toolInstance.Prepare()
			generalConfig := p.prepareConfig(p.parentPipeline,config)
			toolInstanceFlowConfig , err := executor.Run(toolInstance,generalConfig)
			if err != nil {
				executor.Log(fmt.Sprintf("Received Error : %s",err.Error()))

			}else{
				err = p.contextManager.SaveState(toolKey,toolInstanceFlowConfig.GetAsMap())
				if err != nil {
					fmt.Println(fmt.Sprintf("Received Error: %s",err.Error()))
					return
				}

			}
		}else{
			//it is a nested pipeline
			nestedPipelineExecutor := PipelineExecutor{}
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
		RunChildren:
		// Check children
		if vertex.Children.Size() > 0 {
			// Run those children
			for _ , child := range vertex.Children.Values() {
				childFlow := child.(*dag.Vertex)
				p.waitGroup.Add(1)
				p.executeSingleVertex(b,config,childFlow)
			}

		}


	}else{
		//Spawn the current step until all other dependencies are run successfully
		fmt.Println(fmt.Sprintf("Spawning Tool (%s) until dependencies finish execution....",currentFlow.Name))
		p.waitQueue <- vertex
	}


}
func (p *PipelineExecutor) runLocally(b *pipelines.BioPipeline, config models.FlowConfig) error {
	fmt.Println(fmt.Sprintf("Running Pipeline (%s) Locally....",b.Name))
	//Create a Directed Acyclic Graph of the current pipeline
	graph , err := pipelines.CreateGraph(b)
	if err != nil {
		fmt.Println(fmt.Sprintf("Received Error : %s",err.Error()))
		return nil
	}
	parents := graph.SourceVertices()
	//p.waitGroup.Add(graph.Order())
	for _ , parent := range parents{
		//Run each parent individually.
		p.waitGroup.Add(1)
		go p.executeSingleVertex(b,config,parent)
	}
	p.waitGroup.Wait()
	p.stopChan <- nil
	return nil
}
func (p *PipelineExecutor) prepareConfig(b *pipelines.BioPipeline,config models.FlowConfig) models.FlowConfig {
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

func (p *PipelineExecutor) runOnCluster(b *pipelines.BioPipeline, config models.FlowConfig) error {
	return p.runLocally(b,config)
}

func (p *PipelineExecutor) AddTransform(transformCall TransformCall) {
	p.transformations = append(p.transformations,transformCall)
}
func (p *PipelineExecutor) ClearTransformations() bool {
	// Clear all Transformations
	p.transformations = p.transformations[:0]
	return true
}

func (p *PipelineExecutor) Setup(config models.FlowConfig) error {
	p.waitGroup = sync.WaitGroup{}
	p.mutex = &sync.Mutex{}
	p.ticker = time.NewTicker(5 * time.Second)
	p.waitQueue = make(chan *dag.Vertex,5)
	p.stopChan = make(chan interface{},1)
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
func (p *PipelineExecutor) createLogFile(config models.FlowConfig) error {
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

func (p *PipelineExecutor) Log(logs ...interface{}) {
	p.logger.Println(logs...)
}


