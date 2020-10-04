package executors

import (
	"bioflows/models"
	"bioflows/models/pipelines"
)

/*

This function will use "URL" property in Biopipeline definition file and use it to download the tool

 */
var UseUrl TransformCall = func (b *pipelines.BioPipeline,config models.FlowConfig) error{
	//TODO: Download the tool from the given URL if exists and update the current tool definition with it
	return nil
}

var UseBioFlowId TransformCall = func (b *pipelines.BioPipeline,config models.FlowConfig) error {
	//TODO: use BioflowId if found and communicate with Bioflows Hub to download the latest version of this tool using the tool version if found,
	// if the tool version is not found, we will assume that the tool is the latest version
	return nil
}


