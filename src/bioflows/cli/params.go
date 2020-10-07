package cli

import (
	"bioflows/helpers"
	"bioflows/models/pipelines"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
)

func GetRequirementsTableFor(toolPath string) (*simpletable.Table,error){
	table := simpletable.New()
	pipeline := &pipelines.BioPipeline{}
	err := helpers.ReadPipelineFile(pipeline,toolPath)
	if err != nil {
		return nil , err
	}
	graph , err := pipelines.CreateGraph(pipeline)
	if err != nil {
		return nil , err
	}
	successors := graph.SourceVertices()
	if len(successors) <= 0 {
		return nil , errors.New(fmt.Sprintf("BioFlows Definition File: %s is invalid.",pipeline.Name))
	}
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter,Text: "Step #"},
			{Align: simpletable.AlignCenter,Text: "Param. Name (required)"},
			{Align: simpletable.AlignCenter,Text: "Param. Description"},
		},
	}
	for index , parent := range successors {
		parentPipeline := parent.Value.(pipelines.BioPipeline)
		for _ , param := range parentPipeline.Inputs {
			r := []*simpletable.Cell {
				{Align: simpletable.AlignCenter,Text: fmt.Sprintf("%d",index+1)},
				{Align: simpletable.AlignCenter,Text: fmt.Sprintf("%s",param.Name)},
				{Align: simpletable.AlignCenter,Text: fmt.Sprintf("%s",param.GetDescription())},
			}
			table.Body.Cells = append(table.Body.Cells,r)
		}
	}

	return table , nil
}
