package context

import "github.com/hoisie/mustache"

func ParseCommandString(context *BioContext , command string) string {

	processed_command := mustache.Render(command,context)
	return processed_command
}
