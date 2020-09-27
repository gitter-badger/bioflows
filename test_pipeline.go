package main

import (
	"bioflows/engine/pipelines"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main(){
	pipeline := &pipelines.BioPipeline{}
	tool_in, err := os.Open("/home/snouto/projects/bioflows/scripts/pipeline.yaml")

	if err != nil {
		fmt.Printf("There was an error opening the tool file, %v\n",err)
		os.Exit(1)
	}
	mytool_content, err := ioutil.ReadAll(tool_in)
	if err != nil {
		fmt.Printf("Error reading the contents of the tool , %v\n",err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(mytool_content),pipeline)
	if err != nil {
		//fmt.Println("There was a problem unmarshaling the current tool")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(pipeline.ToJson())
}

