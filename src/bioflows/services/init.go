package services

import (
	"bioflows/config"
	"bioflows/kv"
	"github.com/hashicorp/consul/api"
	"strings"
	"fmt"
)

const (
	SERVICES_SECTION_NAME="services"
	SERVICES_ORCHESTRATOR_KEY="type"
	ORCHESTRATION_TYPE_CONSUL="consul"
	ORCHESTRATION_TYPE_ZOOKEEPER = "zookeeper"
)


const (
	SERVICE_NAME = "services/agents/%s"
	SERVICE_HTTP_MODULE_NAME = "services/agents/%s/http"
	SERVICE_TASKS_COUNT_NAME = "services/agents/%s/tasks"
)



type Orchestrator interface {
	Setup(credentials kv.Credentials) error
	FindService(serviceName , tag string , passingOnly bool) ([]*api.ServiceEntry, *api.QueryMeta, error)
	Deregister(id string) error
	Register(name string , address string, port int) error
	Services() (map[string]*api.AgentService,error)
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
