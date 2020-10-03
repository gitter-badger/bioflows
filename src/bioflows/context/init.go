package context

import (
	"fmt"
)

type BioContext struct {
	vars map[string]interface{}
}

func (c *BioContext) init(){
	if c.vars == nil{
		c.vars = make(map[string]interface{})
	}
}

func (c *BioContext) AddVar(key string, value interface{}) bool {
	c.init()
	c.vars[key] = value
	return true
}

func (c *BioContext) HasKey(key string) bool {
	c.init()
	_ , ok := c.vars[key]
	return ok
}

func (c *BioContext) GetKeys() []string{
	keys := make([]string,0)
	for k , _ := range c.vars {
		keys = append(keys,k)
	}
	return keys
}

func (c *BioContext) GetKey(key string) (value interface{}, err error){
	c.init()
	if c.HasKey(key){
		value = c.vars[key]
		err = nil
	}else{
		value = nil
		err = fmt.Errorf("Key not Found")
	}
	return
}




