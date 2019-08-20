package virtualization

import (
	"bioflows/config"
	"bioflows/logs"
	"bytes"
	"strings"
)

const (
	VIRTUALIZATION_SECTION_NAME="virtualization"
	VIRTUALIZATION_KEY_NAME="manager_name"

	DOCKER_VIRTUALIZATION = "docker"
	SINGULARITY_VIRTUALIZATION="singularity"
)

type BioflowContainer struct {
	Names []string
	ID string
	Image string
	ImageID string
}

type VirtualizationManager interface {

	//Returns list of container names that are running on the local node
	ListContainers() []BioflowContainer
	StartContainer(containerName string ,imageName string,commands []string) (*bytes.Buffer , *bytes.Buffer , error)
	StopContainer (containerID string) error
	PullImage(imageURL string )  (*bytes.Buffer , error)
}

var activeManager VirtualizationManager

func init(){

	result, err := config.HasKey(VIRTUALIZATION_SECTION_NAME,VIRTUALIZATION_KEY_NAME)
	if !result && err != nil {

		logs.WriteLog(err.Error())
		return
	}

	val , err := config.GetKeyAsString(VIRTUALIZATION_SECTION_NAME,VIRTUALIZATION_KEY_NAME)
	if err != nil {
		logs.WriteLog(err.Error())
		return
	}

	switch(strings.ToLower(val)){

	case SINGULARITY_VIRTUALIZATION:
		activeManager = &SingularityVirtualizationManager{}
		break;

	case DOCKER_VIRTUALIZATION:
		fallthrough
	default:
		activeManager = &DockerVirtualizationManager{}
		break
	}

}

func GetDefaultVirtualizationManager() VirtualizationManager{
	return activeManager
}
