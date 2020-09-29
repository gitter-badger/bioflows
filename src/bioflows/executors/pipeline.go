package executors

import (
	"bioflows/managers"
	"bioflows/models"
	"bioflows/models/pipelines"
	"fmt"
	"github.com/goombaio/dag"
	"sync"
)

type PipelineExecutor struct {
	contextManager *managers.ContextManager
	planManager *managers.ExecutionPlanManager
	transformations []TransformCall
	waitGroup sync.WaitGroup
}

func (p *PipelineExecutor) IsRemote() bool {
	return p.contextManager.IsRemote()
}

func (p *PipelineExecutor) GetContext() *managers.ContextManager {
	return p.contextManager
}
func (p *PipelineExecutor) Run(b *pipelines.BioPipeline,config models.FlowConfig) models.FlowConfig {
	// Start processing the current pipeline
	PreprocessPipeline(b,config,p.transformations...)
	if p.IsRemote(){
		return p.runOnCluster(b,config)
	}else{
		return p.runLocally(b,config)
	}

}
func (p *PipelineExecutor) runLocally(b *pipelines.BioPipeline, config models.FlowConfig) models.FlowConfig {
	fmt.Println(fmt.Sprintf("Running Pipeline (%s) Locally....",b.Name))
	//Create a Directed Acyclic Graph of the current pipeline
	graph , err := pipelines.CreateGraph(b)
	if err != nil {
		fmt.Println(fmt.Sprintf("Received Error : %s",err.Error()))
		return nil
	}
	parents := graph.SourceVertices()
	p.waitGroup.Add(len(parents))
	finalFlowConfig := models.FlowConfig{}
	var finalError error = nil

	for _ , parent := range parents{

		//Run each parent individually.
		go func(config models.FlowConfig , parent *dag.Vertex) {
			defer func(){
				if r := recover(); r != nil {
					fmt.Println(r)
					return
				}
			}()
			currentFlow := parent.Value.(pipelines.BioPipeline)
			fmt.Println(fmt.Sprintf("Running Tool (%s) Now....",currentFlow.Name))
			if currentFlow.IsTool() {
				executor := ToolExecutor{}
				toolInstance := &models.ToolInstance{
					WorkflowID: b.ID,
					WorkflowName: b.Name,
					Tool:currentFlow.ToTool(),
				}
				toolInstance.Prepare()
				toolInstanceFlowConfig , err := executor.Run(toolInstance,config)
				if err != nil {
					executor.Log(fmt.Sprintf("Received Error : %s",err.Error()))
					finalError = err
				}else{
					finalFlowConfig[toolInstance.ID] = toolInstanceFlowConfig
				}
				p.waitGroup.Done()
				//End of running the tool instance.......
			}else{
				//It is a nested pipeline
			}

		}(config,parent)
	}
	p.waitGroup.Wait()
	return finalFlowConfig
}

func (p *PipelineExecutor) runOnCluster(b *pipelines.BioPipeline, config models.FlowConfig) models.FlowConfig {
	return nil
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
	p.transformations = make([]TransformCall,0)
	p.contextManager = &managers.ContextManager{}
	p.planManager = &managers.ExecutionPlanManager{}
	err := p.contextManager.Setup(config)
	if err != nil {
		return err
	}
	p.planManager.SetContextManager(p.contextManager)
	return p.planManager.Setup(config)
}


