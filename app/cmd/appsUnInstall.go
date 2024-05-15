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
var appsUnInstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall an application",
	Long:  `uninstall an application.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newAppId := args[0]
		available := applications.ListApplications()
		if newApp, isMapContainsKey := available[newAppId]; isMapContainsKey {
			switch newApp.Status {
			case applications.Installed:
				err := applications.UnInstall(newApp)
				if err == nil {
					output.Infof("application [%s] installed.", newAppId)
				} else {
					output.Errorf("could not install application [%s]. [%s] ", newAppId, err)
				}
			case applications.Activated:
				output.Errorf("Cannot uninstall application. application [%s] is activated in current profile.", newAppId)
			case applications.Available:
				output.Errorf("application [%s] is not installed.", newAppId)

			default:
				output.Errorf("application [%s] is already installed. Status of the application is [%s]", newAppId, style.AppStatus(newApp.Status))
			}
		} else {
			output.Errorf("application [%s] is not available for installation.", newAppId)
		}
	},
}

func init() {
	appsCmd.AddCommand(appsUnInstallCmd)

}
