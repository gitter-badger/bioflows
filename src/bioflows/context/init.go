package context

import (
	"bioflows/config"
	"fmt"
)

type BioContext struct {
	Vars map[string]string
}

func (c BioContext) init(){
	if c.Vars == nil{
		c.Vars = make(map[string]string)
	}
}

func (c BioContext) AddVar(key , value string) bool {
	c.init()
	c.Vars[key] = value
	return true
}

func (c BioContext) HasKey(key string) bool {
	c.init()
	_ , ok := c.Vars[key]
	return ok
}

func (c BioContext) GetKey(key string) (value string , err error){
	c.init()
	if c.HasKey(key){
		value = c.Vars[key]
		err = nil
	}else{
		value = ""
		err = fmt.Errorf("Key not Found")
	}
	return
}



func NewContext() *BioContext {
	return &BioContext{}
}

func GetDefaultContext() (*BioContext , error) {
	default_context := &BioContext{}
	cfg , err := config.GetConfig()
	if err != nil {
		return nil , err
	}
	for _ , section := range(cfg.Sections()){

		for _ , key := range(section.Keys()){

			default_context.AddVar(key.Name(),key.Value())

		}

	}
	return default_context , nil
}

