package helpers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func DownloadBioFlowFile(b interface{}, path string) error {
	resp , err := http.Get(path)
	if err != nil {
		return err
	}
	data , err := ioutil.ReadAll(resp.Body)
	return yaml.Unmarshal(data,b)
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
