package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type BioWorkflow struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Snippet string `json:"snippet,omitempty"`
	Description string `json:"description,omitempty"`
	Discussions []string `json:"discussions,omitempty"`
	References []Reference `json:"references,omitempty"`
	Maintainers []Maintainer `json:"maintainers,omitempty"`
	Tools []ToolInstance `json:"tools"`
	Links []BioLink `json:"links"`
	Status int `json:"status"`
}

func (workflow BioWorkflow) GetIdentifier() string{
	return fmt.Sprintf("%s-%s",workflow.Name,workflow.ID)
}

func (workflow BioWorkflow) ToJson() string {
	var buffer *bytes.Buffer = &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if encoder != nil{
		encoder.Encode(workflow)
	}
	return buffer.String()
}

func (workflow BioWorkflow) PrepareInstallations() [][]string {
	installations := make([][]string,0)
	for _, instances := range workflow.Tools {
		toolInstallations := make([]string,0)
		for _,instance := range instances.PrepareInstallations(){

			toolInstallations = append(toolInstallations,instance...)
		}
		installations = append(installations,toolInstallations)

	}
	return installations
}