package cli

import (
	"bioflows/config"
	"bioflows/exec"
	"bioflows/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func RunTool(toolPath string,workflowId string , workflowName string,outputDir string) error{
	tool := &models.Tool{}
	tool_in, err := os.Open(toolPath)
	if err != nil {
		fmt.Printf("There was an error opening the tool file, %v\n",err)
		os.Exit(1)
	}
	mytool_content, err := ioutil.ReadAll(tool_in)
	if err != nil {
		fmt.Printf("Error reading the contents of the tool , %v\n",err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(mytool_content),tool)
	if err != nil {
		//fmt.Println("There was a problem unmarshaling the current tool")
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Executing the current tool.")
	executor := exec.ToolExecutor{}
	workflowConfig := models.FlowConfig{}
	workflowConfig[config.WF_INSTANCE_OUTDIR] = outputDir
	tool_name := tool.Name
	if len(tool_name) <= 0 {
		tool_name = workflowName
	}
	err = executor.Run(&models.ToolInstance{WorkflowID: workflowId,Name: workflowName ,WorkflowName: workflowName,Tool:tool},workflowConfig)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
