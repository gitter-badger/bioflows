package models

import "strings"

type Scriptable string

func (s Scriptable) GetCode() string {
	return string(s)
}

type Scripts struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	Code string `json:"code" yaml:"code"`
	Order string `json:"order,omitempty" yaml:"order,omitempty"`
}

func (s Scripts) IsBefore() bool {
	if len(s.Order) == 0 {
		return true
	}else if strings.ToLower(s.Order) == "before"{
		return true
	}
	return false
}

func (s Scripts) IsAfter() bool {
	if len(s.Order) == 0 {
		return false
	}else if len(s.Order) > 0 && strings.ToLower(s.Order) == "after"{
		return true
	}
	return false
}
