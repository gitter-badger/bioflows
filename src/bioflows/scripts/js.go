package scripts

import (
	"bioflows/models"
	"github.com/dop251/goja"
)

type ScriptManager interface {
	Prepare(toolInstance *models.ToolInstance)
	RunBefore(toolInstance *models.ToolInstance,config map[string]interface{}) error
	RunAfter(toolInstance *models.ToolInstance,config map[string]interface{}) error
}

type JSScriptManager struct {

}
func (manager *JSScriptManager) Prepare(toolInstance *models.ToolInstance) {


}
func (manager *JSScriptManager) RunBefore(toolInstance *models.ToolInstance,config map[string]interface{}) error {
	vm := goja.New()
	manager.Prepare(toolInstance)
	config["command"] = toolInstance.Command.ToString()
	vm.Set("self",config)

	beforeScripts := make([]models.Script,0)
	for idx , script := range toolInstance.Scripts {
		if script.Before {
			if script.Order <= 0 {
				script.Order = idx + 1
			}
			beforeScripts = append(beforeScripts,script)
		}
	}
	//sort the scripts according to the assigned orders
	for _ , beforeScript := range beforeScripts{
		_ , err := vm.RunString(beforeScript.Code.ToString())
		if err != nil {
			return  err
		}
	}
	return nil
}

func (manager *JSScriptManager) RunAfter(toolInstance *models.ToolInstance,config map[string]interface{}) error {
	return nil
}
