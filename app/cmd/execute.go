/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/shell"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "execute",
	Aliases: []string{"call"},
	Short:   "executes a batch script.",
	Long: `executes a batch script. 
	    execute <scriptname>`,
	Args:    cobra.ExactArgs(1),
	Example: "execute mybatch.script",
	Run: func(cmd *cobra.Command, args []string) {
		scriptname := args[0]
		shell.ExecuteScript(scriptname)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
