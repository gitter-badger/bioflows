package cmd

import (
	"bioflows/cli"
	"errors"
	"github.com/spf13/cobra"
)

var workflowRunCmd = &cobra.Command{
	Use:"run [pipeline file .bp]",
	Short: "",
	Long:"",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1{
			return errors.New("Please provide a pipeline file to run.")
		}
		if len(DataDir) < 1{
			return errors.New("Data Directory Flag is required.")
		}
		if len(OutputDir) < 1 {
			return errors.New("Output Directory Flag is required.")
		}
		toolPath := args[0]
		return cli.RunPipeline(cfgFile,toolPath,OutputDir,DataDir)
	},
}

func init(){
	workflowRunCmd.MarkFlagRequired(OutputDir)
	workflowRunCmd.MarkFlagRequired(DataDir)
	WorkflowCmd.AddCommand(workflowRunCmd)
}