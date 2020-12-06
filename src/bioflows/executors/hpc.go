package executors

import (
	"bioflows/models"
	"github.com/dgruber/drmaa"
	"github.com/dgruber/drmaa2os/pkg/jobtracker/libdrmaa"
)

func SubmitToolAsJob(tool *models.ToolInstance,toolCommand string) (interface{},error){
	jt := drmaa.JobTemplate{}
	ojt , err := models.ConvertJTToDrmaa2(tool.JobTemplate)
	if err != nil {
		return nil , err
	}
	err =libdrmaa.ConvertDRMAA2JobTemplateToDRMAAJobTemplate(*ojt,&jt)
	if err != nil {
		return nil , err
	}
	//create new session
	s , err := drmaa.MakeSession()
	if err != nil {
		return nil , err
	}
	defer s.Exit()
	defer s.DeleteJobTemplate(&jt)
	jt.SetRemoteCommand("bash")
	if toolCommand != "" || len(toolCommand) > 0 {
		jt.SetArgs([]string{"-c",toolCommand})
	}else{
		jt.SetArgs([]string{"-c",tool.JobTemplate.RemoteCommand})
	}
	jobID , err := s.RunJob(&jt)
	if err != nil {
		return nil , err
	}
	jinfo , err := s.Wait(jobID,drmaa.TimeoutWaitForever)
	if err != nil {
		return nil , err
	}
	return jinfo , nil
}
