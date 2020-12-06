package models

import (
	"github.com/dgruber/drmaa2interface"
	"time"
)

type Extension struct {
	ExtensionList map[string]string // stores the extension requests as string
}
type JobTemplate struct {
	Extension `json:"-" xml:"-" yaml:"-"`
	RemoteCommand     string            `json:"remoteCommand,omitempty" yaml:"remoteCommand,omitempty"`
	Args              []string          `json:"args,omitempty" yaml:"args,omitempty"`
	SubmitAsHold      bool              `json:"submitAsHold,omitempty" yaml:"submitAsHold,omitempty"`
	ReRunnable        bool              `json:"reRunnable,omitempty" yaml:"reRunnable,omitempty"`
	JobEnvironment    map[string]string `json:"jobEnvironment,omitempty" yaml:"jobEnvironment,omitempty"`
	WorkingDirectory  string            `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`
	JobCategory       string            `json:"jobCategory,omitempty" yaml:"jobCategory,omitempty"`
	Email             []string          `json:"email,omitempty" yaml:"email,omitempty"`
	EmailOnStarted    bool              `json:"emailOnStarted,omitempty" yaml:"emailOnStarted,omitempty"`
	EmailOnTerminated bool              `json:"emailOnTerminated,omitempty" yaml:"emailOnTerminated,omitempty"`
	JobName           string            `json:"jobName,omitempty" yaml:"jobName,omitempty"`
	InputPath         string            `json:"inputPath,omitempty" yaml:"inputPath,omitempty"`
	OutputPath        string            `json:"outputPath,omitempty" yaml:"outputPath,omitempty"`
	ErrorPath         string            `json:"errorPath,omitempty" yaml:"errorPath,omitempty"`
	JoinFiles         bool              `json:"joinFiles,omitempty" yaml:"joinFiles,omitempty"`
	ReservationID     string            `json:"reservationID,omitempty" yaml:"reservationID,omitempty"`
	QueueName         string            `json:"queueName,omitempty" yaml:"queueName,omitempty"`
	MinSlots          int64             `json:"minSlots,omitempty" yaml:"minSlots,omitempty"`
	MaxSlots          int64             `json:"maxSlots,omitempty" yaml:"maxSlots,omitempty"`
	Priority          int64             `json:"priority,omitempty" yaml:"priority,omitempty"`
	CandidateMachines []string          `json:"candidateMachines,omitempty" yaml:"candidateMachines,omitempty"`
	MinPhysMemory     int64             `json:"minPhysMemory,omitempty" yaml:"minPhysMemory,omitempty"`
	MachineOs         string            `json:"machineOs,omitempty" yaml:"machineOs,omitempty"`
	MachineArch       string            `json:"machineArch,omitempty" yaml:"machineArch,omitempty"`
	StartTime         time.Time         `json:"startTime,omitempty" yaml:"startTime,omitempty"`
	DeadlineTime      time.Time         `json:"deadlineTime,omitempty" yaml:"deadlineTime,omitempty"`
	StageInFiles      map[string]string `json:"stageInFiles,omitempty" yaml:"stageInFiles,omitempty"`
	StageOutFiles     map[string]string `json:"stageOutFiles,omitempty" yaml:"stageOutFiles,omitempty"`
	ResourceLimits    map[string]string `json:"resourceLimits,omitempty" yaml:"resourceLimits,omitempty"`
	AccountingID      string            `json:"accountingString,omitempty" yaml:"accountingString,omitempty"`
}

func ConvertJTToDrmaa2(t *JobTemplate) (*drmaa2interface.JobTemplate,error) {
	jt := &drmaa2interface.JobTemplate{
		Extensible:        nil,
		Extension:         drmaa2interface.Extension{
			ExtensionList: t.ExtensionList,
		},
		RemoteCommand:     t.RemoteCommand,
		Args:              t.Args,
		SubmitAsHold:      t.SubmitAsHold,
		ReRunnable:        t.ReRunnable,
		JobEnvironment:    t.JobEnvironment,
		WorkingDirectory:  t.WorkingDirectory,
		JobCategory:       t.JobCategory,
		Email:             t.Email,
		EmailOnStarted:    t.EmailOnStarted,
		EmailOnTerminated: t.EmailOnTerminated,
		JobName:           t.JobName,
		InputPath:         t.InputPath,
		OutputPath:        t.OutputPath,
		ErrorPath:         t.ErrorPath,
		JoinFiles:         t.JoinFiles,
		ReservationID:     t.ReservationID,
		QueueName:         t.QueueName,
		MinSlots:          t.MinSlots,
		MaxSlots:          t.MaxSlots,
		Priority:          t.Priority,
		CandidateMachines: t.CandidateMachines,
		MinPhysMemory:     t.MinPhysMemory,
		MachineOs:         t.MachineOs,
		MachineArch:       t.MachineArch,
		StartTime:         t.StartTime,
		DeadlineTime:      t.DeadlineTime,
		StageInFiles:      t.StageInFiles,
		StageOutFiles:     t.StageOutFiles,
		ResourceLimits:    t.ResourceLimits,
		AccountingID:      t.AccountingID,
	}
	return jt , nil
}
