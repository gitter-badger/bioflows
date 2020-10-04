package cmd

import (
	"bioflows/cli"
	"fmt"
	"github.com/spf13/cobra"
)

var workflowRunCmd = &cobra.Command{
	Use:"run [pipeline file .bp]",
	Short: "",
	Long:"",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("The command requires some parameters. Please provide these paramters..")
		}
		toolPath := args[0]
		return cli.RunPipeline(cfgFile,toolPath,WorkflowId,WorkflowName,OutputDir)
	},
}

func init(){
	workflowRunCmd.Flags().StringVarP(&WorkflowId,"workflowId","i","","Assign a unique Identifier for the workflow in order to distinguish the workflow later.")
	workflowRunCmd.Flags().StringVarP(&WorkflowName,"workflowName","n","myworkflow","Assign a human readable identifier for the current workflow.")
	workflowRunCmd.MarkFlagRequired("workflowId")
	workflowRunCmd.MarkFlagRequired("workflowName")
	WorkflowCmd.AddCommand(workflowRunCmd)
}