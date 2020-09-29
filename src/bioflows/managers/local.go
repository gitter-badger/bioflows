package managers

import (
	"bioflows/context"
	"fmt"
)

type LocalStateManager struct {

	context *context.BioContext

}
func (c *LocalStateManager) GetStateByID(stepId string) (interface{},error){
	if c.context.HasKey(stepId){
		return c.context.GetKey(stepId)
	}
	return nil , nil
}
func (c *LocalStateManager) SetStateByID(stepId string,config interface{}) error {
	r := c.context.AddVar(stepId,config)
	if !r{
		return fmt.Errorf("Unable to add the state for the given stepId")
	}
	return nil
}

func (c *LocalStateManager) Setup(config map[string]interface{}) error {
	c.context = &context.BioContext{}
	return nil
}
