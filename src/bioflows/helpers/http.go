package helpers

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)


func GetFileDetails(details *FileDetails, uri string) error{
	u , err := url.Parse(uri)
	if err != nil {
		return err
	}
	details.Scheme = u.Scheme
	if strings.EqualFold(u.Scheme,"file") {
		details.Base = u.Path[1:]
		details.Local = true
		details.FileName = filepath.Base(uri)
	}else if strings.EqualFold(u.Scheme,"http") || strings.EqualFold(u.Scheme,"https") {
		details.Base = strings.Replace(uri,filepath.Base(uri),"",1)
		details.Local = false
		details.FileName = filepath.Base(uri)
	} else if u.Scheme == "" {
		details.Local = true
		details.Base = strings.Replace(uri,filepath.Base(uri),"",1)
		details.FileName = filepath.Base(uri)
	} else{
		return errors.New("unsupported file scheme used")
	}
	return nil

}




func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == ""  {
		return false
	}
	return true
}

func GetProtocolScheme(toTest string) (scheme string, err error) {
	u ,err := url.Parse(toTest)
	if err != nil {
		scheme = ""
		err = nil
	}
	scheme = u.Scheme
	err = nil
	return
}

func DownloadBioFlowFile(b interface{}, path string) error {
	resp , err := http.Get(path)
	if err != nil {
		return err
	}
	data , err := ioutil.ReadAll(resp.Body)
	return yaml.Unmarshal(data,b)
}
func DownloadRemoteFile(path string) ([]byte,error) {
	resp , err := http.Get(path)
	if err != nil {
		return nil , err
	}
	return ioutil.ReadAll(resp.Body)
}
func ReadLocalBioFlowFile(b interface{}, path string) error {
	tool_in , err := os.Open(path)
	if err != nil {
		return err
	}
	tool_data , err := ioutil.ReadAll(tool_in)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(tool_data,b)
}

func ReadPipelineFile(pipeline interface{} , pipelineFile string) error{
	if IsValidUrl(pipelineFile){
		return DownloadBioFlowFile(pipeline,pipelineFile)
	}else{
		pipeline_in,err := os.Open(pipelineFile)
		if err != nil {
			fmt.Printf("There was an error opening the tool File: %s",err.Error())
			return err
		}
		mypipeline_contents , err := ioutil.ReadAll(pipeline_in)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s",err.Error()))
			return err
		}
		err = yaml.Unmarshal([]byte(mypipeline_contents),pipeline)
		if err != nil {
			fmt.Printf("Error: %s",err.Error())
			return err
		}
		return err
	}
}
