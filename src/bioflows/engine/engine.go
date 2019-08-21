package engine

import "bioflows/models"

type BioFlowEngine struct {

}


func (e *BioFlowEngine) RegisterSelf() error {
	return nil

}

func (e *BioFlowEngine) RunWorkflowFromJSON(workflowJson string) error {

	return nil

}

func (e *BioFlowEngine) RunWorkflowInstance(workflow models.BioWorkflow) error{
	return nil
}

func (e *BioFlowEngine) RunToolInstanceFromJson(toolInstanceJson string) error{
	return nil
}

func (e *BioFlowEngine) RunToolInstance(instance models.ToolInstance) error{
	return nil
}





