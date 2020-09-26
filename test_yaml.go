package main

import (
	"bioflows/config"
	"bioflows/exec"
	"fmt"
	"gopkg.in/yaml.v2"
	"bioflows/models"
	"io/ioutil"
	"os"
)

func main(){

	tool := &models.Tool{}
	tool_in, err := os.Open("/home/snouto/workflows/ls.bt")
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
		return
	}

	executor := exec.ToolExecutor{}
	workflowConfig := models.FlowConfig{}
	workflowConfig[config.WF_INSTANCE_OUTDIR] = "/home/snouto/workflows"
	_, err = executor.Run(&models.ToolInstance{WorkflowID: "myworkflowId",Name: "mytool",WorkflowName: "MyworkflowName",Tool:tool},workflowConfig)
	if err != nil {
		fmt.Println(err)
	}


//
//	fmt.Println(fmt.Sprintf(`
//	Before : %v,
//	After : %v
//`,tool.Scripts[0].IsBefore(),tool.Scripts[0].IsAfter()))
}
