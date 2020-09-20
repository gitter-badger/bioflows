package exec

import (
	"bioflows/models"
	"bioflows/virtualization"
)

type ToolExecutor struct {
	ToolInstance *models.ToolInstance
	ContainerManager *virtualization.VirtualizationManager
}
func (e *ToolExecutor) init() error {
	e.ContainerManager = nil
	return nil
}

func (e *ToolExecutor) Run(t *models.ToolInstance,previousConfig models.FlowConfig,additionalConfig models.FlowConfig) error {
	return nil
}

