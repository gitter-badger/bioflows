package cli

import (
	"bioflows/executors"
	"bioflows/helpers"
	"bioflows/models"
	"bioflows/models/pipelines"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/goombaio/dag"
	"strings"
)

func GetRequirementsTableFor(toolPath string) (*simpletable.Table,error){
	table := simpletable.New()
	pipeline := &pipelines.BioPipeline{}
	err := helpers.ReadPipelineFile(pipeline,toolPath)
	executors.PreprocessPipeline(pipeline,nil,nil)
	if err != nil {
		return nil , err
	}
	graph , err := pipelines.CreateGraph(pipeline)
	if err != nil {
		return nil , err
	}
	var successors []*dag.Vertex

	if pipeline.Type == "tool" {

		successors = make([]*dag.Vertex,0)
		oneVertex := &dag.Vertex{
			Value: *pipeline,
		}
		successors = append(successors,oneVertex)
	}else{
		successors = graph.SourceVertices()
	}

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
	TotalInputs := make([]models.Parameter,0)
	for index , parent := range successors {
		parentPipeline := parent.Value.(pipelines.BioPipeline)
		executors.PreprocessPipeline(&parentPipeline,nil,nil)
		if strings.EqualFold(strings.ToLower(parentPipeline.Type),"pipeline") && len(parentPipeline.Steps) > 0 {

			nestedGraph , err  := pipelines.CreateGraph(&parentPipeline)
			if err != nil {
				return nil , errors.New("Unable to create graph for the nested pipeline.")
			}
			nested_successors := nestedGraph.SourceVertices()
			if len(nested_successors) <= 0 {
				return nil , errors.New(fmt.Sprintf("Nested Pipeline (%s) is invalid",parentPipeline.Name))
			}
			for _ , nested := range nested_successors {
				current := nested.Value.(pipelines.BioPipeline)
				executors.PreprocessPipeline(&current,nil,nil)
				TotalInputs = append(TotalInputs,current.Inputs...)
			}
		}else{
			TotalInputs = append(TotalInputs,parentPipeline.Inputs...)
		}
		for _ , param := range TotalInputs {
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
