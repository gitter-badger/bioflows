package tools

import (
	"bioflows/models"
	"bioflows/logs"
	"fmt"
	"bioflows/process"
	"bioflows/context"
)

func InstallTool(newContext *context.BioContext , tool *models.Tool) error {
	logs.WriteLog(fmt.Sprintf("Installing Tool : %s",tool.Name))
	installations := tool.Installations
	if len(installations) > 0 {
		for _ , installation := range(installations){
			executor := process.CommandExecutor{}
			processed_command := context.ParseCommandString(newContext,installation)
			executor.Command = processed_command
			result_error := executor.Run()
			if result_error != nil {
				logs.WriteLog(fmt.Sprintf("Installation Command : %s , For Tool %s has failed",processed_command,tool.Name))
				continue
			}
		}
		logs.WriteLog("Installation Finished")
	}
	return nil
}

func InstallToolFromJson(newContext *context.BioContext , toolJson string) error {
	newTool := JSONToTool(toolJson)
	return InstallTool(newContext , newTool)
}
