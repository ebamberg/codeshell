/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"

	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "manage, install and uninstall applications.",
	Long:  `manage, install and uninstall applications..`,
	Run: func(cmd *cobra.Command, args []string) {
		output.Infoln(cmd.UsageString())
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)

}
