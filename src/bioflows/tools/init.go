package tools

import (
	"encoding/json"
	"bioflows/models"
)


func JSONToTool(toolJson string) *models.Tool {
	newTool := &models.Tool{}
	if err := json.Unmarshal([]byte(toolJson),*newTool); err != nil {
		panic("Unable to convert the current JSON into tool , Please check your JSON")
	}
	return newTool
}

func NewTool() *models.Tool {
	return &models.Tool{}
}


