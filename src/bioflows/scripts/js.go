package scripts

import (
	"bioflows/models"
	"bioflows/scripts/io"
	"github.com/dop251/goja"
)

type ScriptManager interface {
	Prepare(toolInstance *models.ToolInstance)
	RunBefore(script models.Script,config map[string]interface{}) (error)
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
	if manager.toolInstance != nil {
		config["command"] = manager.toolInstance.Command.ToString()
	}else{
		config["command"] = ""
	}
	vm.Set("self",config)
	vm.Set("io",&io.IO{})
	_ , err := vm.RunString(script.Code.ToString())
	if err != nil {
		return  err
	}
	return nil
}

func (manager *JSScriptManager) RunAfter(script models.Script,config map[string]interface{}) error {
	return manager.RunBefore(script,config)
}
