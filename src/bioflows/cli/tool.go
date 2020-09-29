package cli

import (
	"bioflows/config"
	"bioflows/executors"
	"bioflows/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadConfig(cfgFile string) (models.FlowConfig,error) {
	workflowConfig := models.FlowConfig{}
	config_in , err := os.Open(cfgFile)
	config_contents , err := ioutil.ReadAll(config_in)
	err = yaml.Unmarshal(config_contents,&workflowConfig)
	if err != nil {
		fmt.Println(err.Error())
		return nil , err
	}
	return workflowConfig,nil
}

func RunTool(configFile string, toolPath string,workflowId string , workflowName string,outputDir string) error{
	tool := &models.Tool{}
	workflowConfig := models.FlowConfig{}
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
	BfConfig , err := ReadConfig(configFile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	workflowConfig.Fill(BfConfig)
	fmt.Println("Executing the current tool.")
	executor := executors.ToolExecutor{}
	workflowConfig[config.WF_INSTANCE_OUTDIR] = outputDir
	tool_name := tool.Name
	if len(tool_name) <= 0 {
		tool_name = workflowName
	}
	_ ,err = executor.Run(&models.ToolInstance{WorkflowID: workflowId,Name: workflowName ,WorkflowName: workflowName,Tool:tool},workflowConfig)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
