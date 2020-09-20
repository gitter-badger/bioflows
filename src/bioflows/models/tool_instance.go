package models

import (
	"strings"
	"time"
)

type ToolInstance struct {
	*Tool
	Name string `json:"name"`
	WorkflowID string `json:"workflowId"`
	WorkflowName string `json:"workflowName"`
	StartTime time.Duration
	Status int `json:"status"`

}

func (instance ToolInstance) PrepareCommand() []string{
	splitted_command := strings.Split(string(instance.Command)," ")
	if splitted_command[0] == "sudo"{
		splitted_command = splitted_command[1:]
	}
	instanceCommand := make([]string,0)
	instanceCommand = append(instanceCommand,splitted_command...)
	return instanceCommand
}

func (instance ToolInstance) PrepareInstallations() [][]string{

	installations := make([][]string,0)
	for _,installation := range instance.Installations{
		temp_installation := make([]string,0)
		splitted_installation := strings.Split(installation," ")
		if splitted_installation[0] == "sudo" {
			splitted_installation = splitted_installation[1:]
		}
		temp_installation = append(temp_installation,splitted_installation...)
		installations = append(installations,temp_installation)

	}
	return installations

}


func (instance ToolInstance) GetContainerName() string {
	return strings.Join([]string{instance.WorkflowID,instance.ID,instance.Name},"_")
}

