/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/applications"
	"codeshell/output"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all applications",
	Long:  `list all applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		installed := applications.ListInstalledAppications()
		if len(installed) > 0 {
			output.PrintAsTable(installed, func(row any) []string {
				app := row.(applications.Application)
				return []string{app.DisplayName, app.Status.String()}
			})
		} else {
			fmt.Printf("no applications found.")
		}
	},
}

func init() {
	appsCmd.AddCommand(appsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
