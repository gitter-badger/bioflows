package main

import (
	"fmt"
	"github.com/dgruber/drmaa2interface"
	"github.com/dgruber/drmaa2os"
	_ "github.com/dgruber/drmaa2os/pkg/jobtracker/libdrmaa"
)
const sessionID = "jobsession4"
func main(){
	
	sm, err := drmaa2os.NewDefaultSessionManager("testdb.db")
	if err != nil {
		panic(err)
	}
	js , err := sm.CreateJobSession(sessionID,"")

	if err != nil {
		panic(err)
	}
	defer js.Close()
	defer sm.DestroyJobSession(sessionID)
	jt := drmaa2interface.JobTemplate{
		JobName: "mohamedJob",
		RemoteCommand: "sleep",
		Args: []string{"60"},
	}
	job , err := js.RunJob(jt)
	if err != nil {
		panic(err)
	}

	job.WaitTerminated(drmaa2interface.InfiniteTime)
	jobinfo , err := job.GetJobInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s\n", jobinfo.ID)
	fmt.Printf("State: %s\n", jobinfo.State)
	fmt.Printf("SubState: %s\n", jobinfo.SubState)
	fmt.Printf("Annotation: %s\n", jobinfo.Annotation)
	fmt.Printf("ExitStatus: %d\n", jobinfo.ExitStatus)
	fmt.Printf("TerminatingSignal: %s\n", jobinfo.TerminatingSignal)
	fmt.Printf("AllocatedMachines: %v\n", jobinfo.AllocatedMachines)
	fmt.Printf("SubmissionMachine: %s\n", jobinfo.SubmissionMachine)
	fmt.Printf("JobOwner: %s\n", jobinfo.JobOwner)
	fmt.Printf("Slots: %d\n", jobinfo.Slots)
	fmt.Printf("QueueName: %s\n", jobinfo.QueueName)
	fmt.Printf("WallclockTime: %s\n", jobinfo.WallclockTime)
	fmt.Printf("CPUTime: %d\n", jobinfo.CPUTime)
	fmt.Printf("SubmissionTime: %s\n", jobinfo.SubmissionTime)
	fmt.Printf("DispatchTime: %s\n", jobinfo.DispatchTime)
	fmt.Printf("FinishTime: %s\n", jobinfo.FinishTime)

}
