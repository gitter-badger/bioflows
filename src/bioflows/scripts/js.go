package scripts

import (
	"bioflows/models"
	"github.com/dop251/goja"
)

type ScriptManager interface {
	Prepare(toolInstance *models.ToolInstance)
	RunBefore(script models.Script,config map[string]interface{}) error
	RunAfter(script models.Script,config map[string]interface{}) error
}

type JSScriptManager struct {

	toolInstance *models.ToolInstance

}
func (manager *JSScriptManager) Prepare(toolInstance *models.ToolInstance) {
	manager.toolInstance = toolInstance

}
func (manager *JSScriptManager) RunBefore(script models.Script,config map[string]interface{}) error {
	vm := goja.New()
	config["command"] = manager.toolInstance.Command.ToString()
	vm.Set("self",config)
	_ , err := vm.RunString(script.Code.ToString())
	if err != nil {
		return  err
	}
	return nil
}

func (manager *JSScriptManager) RunAfter(script models.Script,config map[string]interface{}) error {
	return nil
}
