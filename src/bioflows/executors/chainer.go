package executors

import (
	"bioflows/models"
	"bioflows/models/pipelines"
)

type TransformCall func (b *pipelines.BioPipeline,config models.FlowConfig)

func PreprocessPipeline(b *pipelines.BioPipeline,config models.FlowConfig, transforms ...TransformCall)  {

	if len(transforms) <= 0{
		return
	}
	for _ , transform := range transforms {
		transform(b,config)
	}
}
