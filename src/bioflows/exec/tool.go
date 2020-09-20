package exec

import (
	"bioflows/config"
	"bioflows/models"
	"bioflows/process"
	"bioflows/virtualization"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ToolExecutor struct {
	ToolInstance *models.ToolInstance
	ContainerManager *virtualization.VirtualizationManager
	toolLogger *log.Logger
	flowConfig models.FlowConfig
}
func (e *ToolExecutor) GetToolOutputDir() (string,error) {
	workflowOutputDir , ok := e.flowConfig[config.WF_INSTANCE_OUTDIR]
	if !ok {
		return "" , fmt.Errorf("Unable to get the Tool/Workflow Output Directory")
	}
	toolOutputDir := strings.Join([]string{e.ToolInstance.Name,e.ToolInstance.ID},"_")
	toolOutputDir = strings.Join([]string{fmt.Sprintf("%v",workflowOutputDir),toolOutputDir},"/")
	return toolOutputDir , nil
}
func (e *ToolExecutor) CreateOutputFile(name string,ext string) (string,error) {

	outputFile := strings.Join([]string{e.ToolInstance.Name,e.ToolInstance.BioflowId,name},"_")
	outputFile = strings.Join([]string{outputFile,ext},".")
	toolOutputDir , err := e.GetToolOutputDir()
	if err != nil {
		return "" , err
	}
	os.Mkdir(toolOutputDir,config.FILE_MODE_WRITABLE_PERM)
	outputFile = strings.Join([]string{toolOutputDir,outputFile},"/")
	return outputFile , nil

}
func (e *ToolExecutor) init(flowConfig models.FlowConfig) error {
	e.ContainerManager = nil
	e.flowConfig = flowConfig
	// initialize the tool logger
	logFileName , err := e.CreateOutputFile("logs","logs")
	if err != nil {
		return err
	}
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
	e.ToolInstance = t
	err := e.init(workflowConfig)
	if err != nil {
		return err
	}
	executor := &process.CommandExecutor{Command: e.ToolInstance.Command.ToString()}
	toolErr  := executor.Run()
	//Create output file for the output of this tool
	toolOutputFile , err := e.CreateOutputFile("stdout","out")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(toolOutputFile,executor.GetOutput().Bytes(),config.FILE_MODE_WRITABLE_PERM)
	if err != nil {
		return err
	}
	//Create err file for this tool
	toolErrFile , err := e.CreateOutputFile("stderr","err")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(toolErrFile,executor.GetError().Bytes(),config.FILE_MODE_WRITABLE_PERM)
	if err != nil {
		return err
	}
	e.Log(fmt.Sprintf("Tool: %s has finished.",e.ToolInstance.Name))
	return toolErr
}

