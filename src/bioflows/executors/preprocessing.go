package executors

import (
	"bioflows/helpers"
	"bioflows/models"
	"bioflows/models/pipelines"
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
			err := helpers.DownloadBioFlowFile(t,b.URL)
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


