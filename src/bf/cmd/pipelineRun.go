package cmd

import (
	"bioflows/cli"
	"errors"
	"github.com/spf13/cobra"
)

var (
	clean bool
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
		return cli.RunPipeline(cfgFile,toolPath,OutputDir,DataDir,ParamsConfig,clean)
	},
}

func init(){
	workflowRunCmd.PersistentFlags().BoolVar(&clean,"clean",false,"This command cleans all metadata associated with this pipeline from the distributed in-memory Key/Value Store, " +
		"in case you are running in a distributed mode. this command has no effect if you are running in a local mode.")
	workflowRunCmd.MarkFlagRequired(OutputDir)
	workflowRunCmd.MarkFlagRequired(DataDir)
	WorkflowCmd.AddCommand(workflowRunCmd)
}