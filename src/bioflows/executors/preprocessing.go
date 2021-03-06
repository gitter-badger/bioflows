package executors

import (
	"bioflows/helpers"
	"bioflows/models"
	"bioflows/models/pipelines"
	"path/filepath"
)

/*

This function will use "URL" property in Biopipeline definition file and use it to download the tool

 */

var UseUrl TransformCall = func (b *pipelines.BioPipeline,config models.FlowConfig) error{
	if len(b.URL) > 0 {
		//check to see the pipeline
		if helpers.IsValidUrl(b.URL) {
			// That means it is a valid URL, so download the file
			t := &pipelines.BioPipeline{}
			fileDetails := &helpers.FileDetails{}
			err := helpers.GetFileDetails(fileDetails,b.URL)
			if err != nil {
				return err
			}
			currPath := filepath.Join(config["bf_tool_basepath"].(string),fileDetails.FileName)
			if fileDetails.Local {
				// that means it is a local file , so read it
				err = helpers.ReadLocalBioFlowFile(t,currPath)
			}else{
				// It could be a remote http/https command, so download the file

				err = helpers.DownloadBioFlowFile(t,currPath)
			}
			if err != nil {
				return err
			}
			err = pipelines.Clone(b,t,config)
			return err
		}
	}
	return nil
}

var UseBioFlowId TransformCall = func (b *pipelines.BioPipeline,config models.FlowConfig) error {
	t := &pipelines.BioPipeline{}
	helpers.DownloadFromBioFlowsHub(t,b.BioflowId,b.Version)
	return pipelines.Clone(b,t,config)
}


