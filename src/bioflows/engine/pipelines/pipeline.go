package pipelines

import (
	"bioflows/helpers/id"
	"bioflows/models"
	"encoding/json"
	"fmt"
	"strings"
)

type BioPipeline struct {
	Type        string   `json:"type,omitempty" yaml:"type,omitempty"`
	Depends     string   `json:"depends,omitempty" yaml:"depends,omitempty"`
	ID          string   `json:"id,omitempty" yaml:"id,omitempty"`
	Order       int      `json:"order,omitempty" yaml:"order,omitempty"`
	BioflowId   string   `json:"bioflowId,omitempty" yaml:"bioflowId,omitempty"`
	URL         string   `json:"url,omitempty" yaml:"url,omitempty"`
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Discussions []string `json:"discussions,omitempty" yaml:"discussions,omitempty"`
	Website     string   `json:"website,omitempty" yaml:"website,omitempty"`
	Version     string   `json:"version,omitempty" yaml:"version,omitempty"`
	Icon        string   `json:"icon,omitempty" yaml:"icon,omitempty"`
	Shadow       bool                `json:"shadow,omitempty" yaml:"shadow,omitempty"`
	Maintainer   *models.Maintainer  `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`
	References   []models.Reference  `json:"references,omitempty" yaml:"references,omitempty"`
	Inputs       []models.Parameter  `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Config       []models.Parameter  `json:"config,omitempty" yaml:"config,omitempty"`
	Outputs      []models.Parameter  `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Command      models.Scriptable   `json:"command" yaml:"command"`
	Dependencies []string            `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Deprecated   bool                `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Conditions   []models.Scriptable `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Steps        []BioPipeline       `json:"steps,omitempty" yaml:"steps,omitempty"`
}

func (instance *BioPipeline) Prepare(){
	//Generate random unique ID if the instance Tool ID is not set
	if len(instance.ID) <= 0 {
		instance.ID , _ = id.NewID()
	}
	//if the tool name is not set, then use the tool ID
	if len(instance.Name) <= 0{
		instance.Name = instance.ID
	}else{
		// If the tool name exists, replace whitespace with underscores
		instance.Name = strings.ReplaceAll(instance.Name," ","_")
	}
	//If the tool name is set , use that as the tool instance name
	instance.Name = instance.Name
}

func (p BioPipeline) ToTool() *models.Tool {
	t := &models.Tool{}
	t.Type = p.Type
	t.Depends = p.Depends
	t.URL = p.URL
	t.ID = p.ID
	t.Order = p.Order
	t.BioflowId = p.BioflowId
	t.Name = p.BioflowId
	t.Description = p.Description
	t.Discussions = make([]string, len(p.Discussions))
	copy(t.Discussions, p.Discussions)
	t.Website = p.Website
	t.Version = p.Version
	t.Icon = p.Icon
	t.Shadow = p.Shadow
	t.Maintainer = p.Maintainer
	t.References = make([]models.Reference, len(p.References))
	copy(t.References, p.References)
	t.Inputs = make([]models.Parameter, len(p.Inputs))
	copy(t.Inputs, p.Inputs)
	t.Config = make([]models.Parameter, len(p.Config))
	copy(t.Config, p.Config)
	t.Outputs = make([]models.Parameter, len(p.Outputs))
	copy(t.Outputs, p.Outputs)
	t.Command = p.Command
	t.Dependencies = make([]string, len(p.Dependencies))
	copy(t.Dependencies, p.Dependencies)
	t.Deprecated = p.Deprecated
	t.Conditions = make([]models.Scriptable, len(p.Conditions))
	return t
}

func (p BioPipeline) IsTool() bool {
	if len(p.Type) <= 0 {
		return true
	}
	if strings.ToLower(p.Type) == "pipeline" || strings.ToLower(p.Type) == "workflow" {
		return false
	}
	return true
}

func (p BioPipeline) IsPipeline() bool {
	return !p.IsTool()
}

func (p *BioPipeline) ToJson() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(fmt.Errorf("Unable to Marshal current BioPipeline into JSON"))
	}
	return string(bytes)
}
