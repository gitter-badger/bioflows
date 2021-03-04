package models

import "encoding/json"

type Tool struct {
	Type         string        `json:"type,omitempty" yaml:"type,omitempty"`
	URL          string        `json:"url,omitempty" yaml:"url,omitempty"`
	Depends      string        `json:"depends,omitempty" yaml:"depends,omitempty"`
	ID           string        `json:"id,omitempty" yaml:"id,omitempty"`
	Order        int           `json:"order,omitempty" yaml:"order,omitempty"`
	ImageId      string        `json:"imageId,omitempty"`
	BioflowId    string        `json:"bioflowId,omitempty" yaml:"bioflowId,omitempty"`
	Name         string        `json:"name" yaml:"name"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	Discussions  []string      `json:"discussions,omitempty" yaml:"discussions,omitempty"`
	Website      string        `json:"website,omitempty" yaml:"website,omitempty"`
	Version      string        `json:"version,omitempty" yaml:"version,omitempty"`
	Icon         string        `json:"icon,omitempty" yaml:"icon,omitempty"`
	Shadow       bool          `json:"shadow,omitempty" yaml:"shadow,omitempty"`
	Maintainer   *Maintainer   `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`
	References   []Reference   `json:"references,omitempty" yaml:"references,omitempty"`
	Inputs       []Parameter   `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Config       []Parameter   `json:"config,omitempty" yaml:"config,omitempty"`
	Outputs      []Parameter   `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Command      Scriptable    `json:"command" yaml:"command"`
	Loop bool 	`json:"loop,omitempty" yaml:"loop,omitempty"`
	LoopVar string `json:"loop_var,omitempty" yaml:"loop_var,omitempty"`
	Dependencies []string      `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Deprecated   bool          `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Conditions   []Scriptable  `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Scripts      []Script      `json:"scripts,omitempty" yaml:"scripts,omitempty"`
	Notification *Notification `json:"notification,omitempty" yaml:"notification,omitempty"`
	Caps         *Capabilities `json:"caps,omitempty" yaml:"caps,omitempty"`
	ContainerConfig *ContainerConfig `json:"container,omitempty" yaml:"container,omitempty"`
}

func (t *Tool) ToJson() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic("Unable to Convert the current tool into JSON.")
	}
	return string(bytes)
}
