/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/profiles"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var profilesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all profiles",
	Long:  `list all defined profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := profiles.ListProfiles()
		if len(p) > 0 {
			header := []string{"id", "Profile Name"}
			output.PrintAsTableH(p, header, func(row any) []string {
				profile := row.(profiles.Profile)
				return []string{profile.Id, profile.Displayname}
			})
		} else {
			output.Println("no profiles found.")
		}
	},
}

func init() {
	profilesCmd.AddCommand(profilesListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
