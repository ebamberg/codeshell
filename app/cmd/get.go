/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get the value of an environment variable",
	Long: `get the value of an environment variable. For example:
			get myVariable.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output.Println(utils.GetEnvVariable(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
