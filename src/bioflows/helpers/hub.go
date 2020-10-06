package helpers

import (
	"bioflows/config"
	"fmt"
	"strings"
)

func DownloadFromBioFlowsHub(tool interface{},bioflowId, version string){
	if len(bioflowId) <= 0 {
		return
	}
	toolHubPath := strings.Join([]string{config.BIOFLOWS_HUB_LOCATION,bioflowId},"/")
	if len(version) > 0 {
		toolHubPath = fmt.Sprintf("%s?version=%s",toolHubPath,version)
	}
	DownloadBioFlowFile(tool,toolHubPath)
}
