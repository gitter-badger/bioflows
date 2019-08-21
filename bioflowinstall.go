package main

import (
	"os"
	"fmt"
	"strings"
	"bioflows/process"
)

func main(){
	args := os.Args

	if len(args) <= 1{
		fmt.Println("You have to supply the commands to install")
		return
	}
	WorkflowId := args[1]
	fmt.Println("Downloading Workflow : ",WorkflowId)

}