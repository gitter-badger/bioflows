package models

import "encoding/json"

/*
   Task is a struct which is used when a command line client submits a pipeline to a running bioflow node ,
	It can be identified by unique Task Id, it carries also byte slice containing the Json representation of the Tool/pipeline,
 	It also contains byte slice representation of configuration parameters (FlowConfig)
 */
const (
	TASK_STATUS_PENDING = iota
	TASK_STATUS_RUNNING
	TASK_STATUS_FAILED
	TASK_STATUS_FINISHED
)
type Task struct {
	StatusId int `json:"statusId,omitempty" yaml:"statusId,omitempty"`
	Retries int `json:"retries,omitempty" yaml:"retries,omitempty"`
	NodeId string `json:"nodeId,omitempty" yaml:"nodeId,omitempty"`
	TaskId string `json:"taskId" yaml:"taskId"`
	Task []byte `json:"task,omitempty" yaml:"task,omitempty"`
	Config []byte `json:"config,omitempty" yaml:"config,omitempty"`
}
func (t *Task) ToJson() (string,error){
	data , err := json.Marshal(t)
	if err != nil {
		return "" , err
	}
	return string(data) , nil
}

func (t *Task) FromJson(data []byte) error {
	return json.Unmarshal(data,t)
}