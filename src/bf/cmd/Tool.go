/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// ToolCmd represents the Tool command
var ToolCmd = &cobra.Command{
	Use:   "Tool",
	Short: "Helper commands to manage and run single BioFlows Tools",
	Long: `Helper commands to manage and run single BioFlows Tools`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	//ToolCmd.PersistentFlags().StringVar(&OutputDir,"output_dir","","Output Directory where the running tool will save data.")
	rootCmd.AddCommand(ToolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ToolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ToolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
