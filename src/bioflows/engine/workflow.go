package engine

import (
	"bioflows/context"
	"bioflows/virtualization"
	"bioflows/filesystem"
	"bioflows/models"
)

type BioFlowManager struct {
	context *context.BioContext
	virtualManager *virtualization.VirtualizationManager
	ioManager *filesystem.FileSystemManager
}

func (manager *BioFlowManager) SetContext(context *context.BioContext){
	manager.context = context
}

func (manager *BioFlowManager) SetVirtualManager(virtual *virtualization.VirtualizationManager){
	manager.virtualManager = virtual
}

func (manager *BioFlowManager) SetIOManager(io *filesystem.FileSystemManager){
	manager.ioManager = io
}

func (manager *BioFlowManager) StartFromId(workflowId string) error {
	return nil
}

func (manager *BioFlowManager) StartFromString(workflowJSON string) error {
	return nil
}

func (manager *BioFlowManager) StartToolFromId(taskId string) error {
	return nil
}

func (manager *BioFlowManager) StartToolFromString(taskJson string) error {
	return nil
}

func (manager *BioFlowManager) StartTool(task models.ToolInstance) error {
	return nil
}



