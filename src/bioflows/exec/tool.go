package exec

import (
	"bioflows/config"
	"bioflows/models"
	"bioflows/virtualization"
	"fmt"
	"log"
	"os"
	"strings"
)

type ToolExecutor struct {
	ToolInstance *models.ToolInstance
	ContainerManager *virtualization.VirtualizationManager
	toolLogger *log.Logger
}
func (e *ToolExecutor) init(flowConfig models.FlowConfig) error {
	e.ContainerManager = nil
	// initialize the tool logger
	workflowOutputDir := flowConfig[config.WF_INSTANCE_OUTDIR]
	logFileName := strings.Join([]string{e.ToolInstance.WorkflowID,e.ToolInstance.BioflowId},"_")
	logFileName = strings.Join([]string{logFileName,"logs"},".")
	logFileName = strings.Join([]string{fmt.Sprintf("%v",workflowOutputDir),logFileName},"/")
	e.toolLogger = &log.Logger{}
	e.toolLogger.SetPrefix(config.BIOFLOWS_NAME)
	file , err := os.Create(logFileName)
	if err != nil {
		fmt.Printf("Can't Create Tool (%s) log file %s",e.ToolInstance.Name, logFileName)
		return err
	}
	e.toolLogger.SetOutput(file)
	return nil
}
func (e *ToolExecutor) Log(logs ...interface{}) {
	e.toolLogger.Print(logs)
}

func (e *ToolExecutor) Run(t *models.ToolInstance, workflowConfig models.FlowConfig) error {
	err := e.init(workflowConfig)
	if err != nil {
		return err
	}

	return nil
}

