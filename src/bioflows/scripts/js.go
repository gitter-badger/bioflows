package scripts

import (
	config2 "bioflows/config"
	"bioflows/helpers"
	"bioflows/models"
	"bioflows/scripts/io"
	"errors"
	"github.com/dop251/goja"
	"io/ioutil"
	"path/filepath"
)

type ScriptManager interface {
	Prepare(toolInstance *models.ToolInstance)
	RunBefore(script models.Script,config map[string]interface{}) (error)
	RunAfter(script models.Script,config map[string]interface{}) error
	getCode(script models.Script , config map[string]interface{}) (string , error)
}

type JSScriptManager struct {

	toolInstance *models.ToolInstance

}
func (manager *JSScriptManager) Prepare(toolInstance *models.ToolInstance) {
	manager.toolInstance = toolInstance

}
func (manager *JSScriptManager) getCode(script models.Script , config map[string]interface{}) (string,error) {
	code := script.Code.ToString()
	if len(code) > 2 {
		return code, nil
	}
	if script.CodeFile != "" && len(script.CodeFile) > 1 {
		details := helpers.FileDetails{}
		err := helpers.GetFileDetails(&details,script.CodeFile)
		if err != nil {
			return "" , err
		}
		tool_basePath := config[config2.WF_BF_TOOL_BASEPATH].(string)
		codeFile_location := filepath.Join(tool_basePath,details.Base)
		if details.Local {
			codefile_data , err := ioutil.ReadFile(codeFile_location)
			if err != nil {
				return "" , err
			}
			return string(codefile_data) , nil
		}else{
			data , err := helpers.DownloadRemoteFile(codeFile_location)
			if err != nil {
				return "" , err
			}
			return string(data) , nil
		}
	}
	return "" , errors.New("invalid script directive. no code found")
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
	code , err := manager.getCode(script,config)
	if err != nil {
		return err
	}
	_ , err = vm.RunString(code)
	if err != nil {
		return  err
	}
	return nil
}

func (manager *JSScriptManager) RunAfter(script models.Script,config map[string]interface{}) error {
	return manager.RunBefore(script,config)
}
