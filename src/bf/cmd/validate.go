package cmd

import (
	"bioflows/cli"
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use: "validate [Tool/Pipeline Path]",
	Short: `validates a given BioFlows tool or pipeline definition file. It checks whether the file is valid and well-formatted or not.
    	The file path could be a Local File System Path or a remote URL.
`,
	Long: `validates a given BioFlows tool or pipeline definition file. It checks whether the file is valid and well-formatted or not.
    	The file path could be a Local File System Path or a remote URL.
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) <= 0 {
			return cmd.Usage()
		}
		filePath := args[0]
		valid, err := cli.ValidateYAML(filePath)
		if err != nil {
			fmt.Println(fmt.Sprintf("%s", err.Error()))
			return err
		}
		if valid {
			fmt.Println("Validate Tool: The Tool is valid.")
		} else {
			fmt.Println("Validate Tool: The tool is not valid.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ValidateCmd)
}
