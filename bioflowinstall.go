package main

import (
	"os"
	"fmt"

)

func main(){
	args := os.Args

	if len(args) <= 1{
		fmt.Println("You have to supply the Workflow Id or tool Id to install")
		return
	}
	WorkflowId := args[1]
	fmt.Println("Downloading Workflow : ",WorkflowId)

}