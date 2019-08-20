package models

import "encoding/json"

type Parameter struct {
	DisplayName string `json:"displayName,omitempty"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Reference struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Website string `json:"website,omitempty"`
}

type Maintainer struct {
	Username string `json:"name,omitempty"`
	FullName string `json:"fullname,omitempty"`
	Email string `json:"email,omitempty"`

}

type Tool struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Discussions []string `json:"discussions,omitempty"`
	Website string `json:"website,omitempty"`
	Version string `json:"version"`
	Icon string `json:"icon,omitempty"`
	Maintainer *Maintainer `json:"maintainer,omitempty"`
	References []Reference `json:"references,omitempty"`
	Inputs []Parameter `json:"inputs,omitempty"`
	Config []Parameter `json:"config,omitempty"`
	Outputs []Parameter `json:"outputs,omitempty"`
	Command string `json:"command"`
	Installations []string `json:"installations,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
}

func (t *Tool) ToJson() string {
	bytes , err := json.Marshal(t)
	if err != nil {
		panic("Unable to Convert the current tool into JSON.")
	}
	return string(bytes)
}
