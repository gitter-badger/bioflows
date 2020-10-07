package models

import (
	"fmt"
)

type FlowConfig map[string]interface{}

func (f *FlowConfig) Fill(config map[string]interface{})  {
	for k , v := range config {
		(*f)[k] = v
	}

}

func (f *FlowConfig) GetAsMap() map[string]interface{}{
	newMap := make(map[string]interface{})
	for k , v := range *f {
		newMap[k] = v
	}
	return newMap
}

type Parameter struct {
	DisplayName string      `json:"displayname,omitempty" yaml:"displayname,omitempty"`
	Name        string      `json:"name" yaml:"name"`
	Type        string      `json:"type,omitempty" yaml:"type,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value       interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

func (p *Parameter) GetParamValue() string {
	return fmt.Sprintf("%v", p.Value)
}
func (p *Parameter) GetDescription() string{
	if len(p.Description) > 0{
		return p.Description
	}else{
		return p.DisplayName
	}
}

type Reference struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Website     string `json:"website,omitempty" yaml:"website,omitempty"`
}

type Maintainer struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	FullName string `json:"fullname,omitempty" yaml:"fullname,omitempty"`
	Email    string `json:"email,omitempty" yaml:"email,omitempty"`
}

type Notification struct {
	To    string `json:"to" yaml:"to"`
	CC    string `json:"cc,omitempty" yaml:"cc,omitempty"`
	Title string `json:"title" yaml:"title"`
	Body  string `json:"body" yaml:"body"`
}

type Capabilities struct {
	CPU int `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory int `json:"memory,omitempty" yaml:"memory,omitempty"`
}

