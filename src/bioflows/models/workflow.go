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
	ImageId string `json:"imageId,omitempty"`
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
