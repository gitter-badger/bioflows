package services

import (
	"bioflows/config"
	"strings"
	"fmt"
)

const (
	SERVICES_SECTION_NAME="services"
	SERVICES_ORCHESTRATOR_KEY="type"
	ORCHESTRATION_TYPE_CONSUL="consul"
	ORCHESTRATION_TYPE_ZOOKEEPER = "zookeeper"
)

type Orchestrator interface {


}
func GetDefaultOrchestrator() (Orchestrator , error){

	val , err := config.GetKeyAsString(SERVICES_SECTION_NAME,SERVICES_ORCHESTRATOR_KEY)
	if err != nil {
		return nil , err
	}
	switch(strings.ToLower(val)){
	case ORCHESTRATION_TYPE_CONSUL:
		return &ConsulOrchestrator{} , nil
	case ORCHESTRATION_TYPE_ZOOKEEPER:
		return &ZooKeeperOrchestrator{},nil
	default:
		return nil , fmt.Errorf("Unknown Orchestrator Used")
	}

}
