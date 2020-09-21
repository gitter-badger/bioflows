package exec

import (
	"bioflows/config"
	"bioflows/expr"
	"bioflows/models"
	"bioflows/process"
	"bioflows/scripts"
	"bioflows/virtualization"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type ToolExecutor struct {
	ToolInstance *models.ToolInstance
	ContainerManager *virtualization.VirtualizationManager
	toolLogger *log.Logger
	flowConfig models.FlowConfig
	exprManager *expr.ExprManager
}
func (e *ToolExecutor) prepareParameters() models.FlowConfig {
	flowConfig := make(models.FlowConfig)
	//Copy all flow configs at the workflow level into the current tool flowconfig
	if len(e.flowConfig) > 0 {
		for k,v := range e.flowConfig{
			flowConfig[k] = v
		}
	}
	if len(e.ToolInstance.Inputs) <= 0 {
		return flowConfig
	}
	inputs := make(map[string]string)
	for _ , param := range e.ToolInstance.Inputs{
		paramValue := e.exprManager.Render(param.GetParamValue(),flowConfig)
		inputs[param.Name] = paramValue
	}
	//Append the processed input parameters into the current flowConfig
	for k , v := range inputs {
		flowConfig[k] = v
	}
	if len(e.ToolInstance.Outputs) <= 0{
		return flowConfig
	}
	//Prepare outputs
	outputs := make(map[string]string)
	for _ , param := range e.ToolInstance.Outputs {
		paramValue := e.exprManager.Render(param.GetParamValue(),flowConfig)
		outputs[param.Name] = paramValue
	}
	for k,v  := range outputs{
		flowConfig[k] = v
	}
	return flowConfig
}
func (e *ToolExecutor) executeBeforeScripts() (map[string]interface{},error) {
	configuration := e.prepareParameters()
	beforeScripts := make([]models.Script,0)
	for idx , script := range e.ToolInstance.Scripts {
		if script.IsBefore() {
			if script.Order <= 0 {
				script.Order = idx + 1
			}
			beforeScripts = append(beforeScripts,script)
		}
	}
	//sort the scripts according to the assigned orders
	sort.Slice(beforeScripts, func(i, j int) bool {

		return beforeScripts[i].Order < beforeScripts[j].Order

	})
	for _ , beforeScript := range beforeScripts {
		var scriptManager scripts.ScriptManager
		switch strings.ToLower(beforeScript.Type) {
		case "js":
			fallthrough
		default:
			scriptManager = &scripts.JSScriptManager{}
			scriptManager.Prepare(e.ToolInstance)
		}
		err := scriptManager.RunBefore(beforeScript,configuration)
		if err != nil {
			return configuration , err
		}
	}
	return configuration , nil
}
func (e *ToolExecutor) executeAfterScripts() (map[string]interface{},error)  {
	configuration := e.prepareParameters()
	afterScripts := make([]models.Script,0)
	for idx , script := range e.ToolInstance.Scripts {
		if script.IsAfter() {
			if script.Order <= 0 {
				script.Order = idx + 1
			}
			afterScripts = append(afterScripts,script)
		}
	}
	//sort the scripts according to the assigned orders
	sort.Slice(afterScripts, func(i, j int) bool {

		return afterScripts[i].Order < afterScripts[j].Order

	})
	for _ , afterScript := range afterScripts {
		var scriptManager scripts.ScriptManager
		switch strings.ToLower(afterScript.Type) {
		case "js":
			fallthrough
		default:
			scriptManager = &scripts.JSScriptManager{}
			scriptManager.Prepare(e.ToolInstance)
		}
		err := scriptManager.RunAfter(afterScript,configuration)
		if err != nil {
			return configuration , err
		}
	}
	return configuration , nil
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
	e.exprManager = &expr.ExprManager{}
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
func (e *ToolExecutor) execute() error {

	//prepare parameters
	toolConfig, err := e.executeBeforeScripts()
	if err != nil {
		return err
	}
	toolCommandStr := fmt.Sprintf("%v",toolConfig["command"])
	toolCommand := e.exprManager.Render(toolCommandStr,toolConfig)
	executor := &process.CommandExecutor{Command: toolCommand}
	executor.Init()
	toolErr  := executor.Run()
	if e.ToolInstance.Shadow{
		return toolErr
	}
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
func (e *ToolExecutor) Run(t *models.ToolInstance, workflowConfig models.FlowConfig) error {
	e.ToolInstance = t
	err := e.init(workflowConfig)
	if err != nil {
		return err
	}
	return e.execute()
}

