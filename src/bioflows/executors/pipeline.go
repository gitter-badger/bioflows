package executors

import "bioflows/managers"

type PipelineExecutor struct {
	contextManager *managers.ContextManager
	planManager *managers.ExecutionPlanManager

}

func (p *PipelineExecutor) IsRemote() bool {
	return p.contextManager.IsRemote()
}

func (p *PipelineExecutor) GetContext() *managers.ContextManager {
	return p.contextManager
}

func (p *PipelineExecutor) Setup(config map[string]interface{}) error {
	p.contextManager = &managers.ContextManager{}
	p.planManager = &managers.ExecutionPlanManager{}
	err := p.planManager.Setup(config)
	if err != nil {
		return err
	}
	return p.contextManager.Setup(config)
}


