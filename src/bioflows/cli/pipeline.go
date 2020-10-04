package cli

import (
	"bioflows/config"
	"bioflows/executors"
	"bioflows/models"
	"bioflows/models/pipelines"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func RunPipeline(configFile,toolPath,workflowId,workflowName,outputDir string) error{
	fmt.Println(fmt.Sprintf("Using Configuration File: %s",configFile))
	pipeline := &pipelines.BioPipeline{}
	workflowConfig := models.FlowConfig{}
	pipeline_in,err := os.Open(toolPath)
	if err != nil {
		fmt.Printf("There was an error opening the tool File: %s",err.Error())
		return err
	}
	mypipeline_contents , err := ioutil.ReadAll(pipeline_in)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s",err.Error()))
		return err
	}
	err = yaml.Unmarshal([]byte(mypipeline_contents),pipeline)
	if err != nil {
		fmt.Printf("Error: %s",err.Error())
		return err
	}
	BfConfig, err := ReadConfig(configFile)
	if err != nil {
		fmt.Printf("Error Reading Config: %s",err.Error())
		return err
	}
	workflowConfig.Fill(BfConfig)
	workflowConfig[config.WF_INSTANCE_OUTDIR] = outputDir
	fmt.Println(fmt.Sprintf("Executing Workflow: %s",pipeline.Name))
	executor := executors.PipelineExecutor{}
	err = executor.Setup(workflowConfig)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s",err.Error()))
		return err
	}
	err =  executor.Run(pipeline,workflowConfig)
	return err
}
