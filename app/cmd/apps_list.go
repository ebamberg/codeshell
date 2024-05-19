/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/applications"
	"codeshell/output"
	"codeshell/style"
	"fmt"

	"github.com/spf13/cobra"
)

var showall bool

// listCmd represents the list command
var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all applications",
	Long:  `list all applications.`,
	Run: func(cmd *cobra.Command, args []string) {

		allFlag := cmd.Flag("all")
		if allFlag.Value.String() == "true" {
			showall = true
		} else {
			showall = false
		}

		var apps []applications.Application

		if showall {
			apps = applications.FlattenMap(applications.ListApplications())
		} else {
			apps = applications.FlattenMap(applications.ListInstalledAppications())
		}

		if len(apps) > 0 {
			header := []string{"Id", "Name", "Version", "Status"}
			output.PrintTidySlice(apps, header, func(row any) []string {
				app := row.(applications.Application)
				return []string{app.Id, app.DisplayName, app.Version, style.AppStatus(app.Status)}
			})
		} else {
			fmt.Printf("no applications found.")
		}
	},
}

func init() {
	appsCmd.AddCommand(appsListCmd)

	appsListCmd.Flags().BoolVarP(&showall, "all", "a", false, "show all application, available and installed")
}
