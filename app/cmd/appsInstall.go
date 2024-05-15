/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/applications"
	"codeshell/output"
	"codeshell/style"

	"github.com/spf13/cobra"
)

// appsInstallCmd represents the appsInstall command
var appsInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newAppId := args[0]
		available := applications.ListApplications()
		if newApp, isMapContainsKey := available[newAppId]; isMapContainsKey {
			if newApp.Status == applications.Available {
				err := applications.Install(newApp)
				if err == nil {
					output.Infof("application [%s] installed.", newAppId)
				} else {
					output.Errorf("could not install application [%s]. ", newAppId, err)
				}
			} else {
				output.Errorf("application [%s] is already installed. Status of the application is [%s]", newAppId, style.AppStatus(newApp.Status))
			}
		} else {
			output.Errorf("application [%s] is not available for installation.", newAppId)
		}
	},
}

func init() {
	appsCmd.AddCommand(appsInstallCmd)

}
