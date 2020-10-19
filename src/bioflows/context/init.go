package context

import (
	"fmt"
	"strings"
)

type BioContext struct {
	vars map[string]interface{}
}

func (c *BioContext) DeleteByKey(key string) bool {
	keys := c.FilterKeys(key)
	if len(keys) > 0 {
		for _ , k := range keys {
			delete(c.vars,k)
		}
	}
	return true
}
func (c *BioContext) FilterKeys(key string) []string{
	tempKeys := make([]string,0)
	for k , _ := range c.vars {
		if strings.HasPrefix(k,key) {
			tempKeys = append(tempKeys,k)
		}
	}
	return tempKeys
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




