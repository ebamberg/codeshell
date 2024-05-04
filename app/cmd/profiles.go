/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "profiles are a set of installed tools",
	Long: `A profile is a set of installed tools with a specific version.
	
	you can define different profiles with a different set of activated tools and versions.
	For example a specific Java version bundled with a specific npm version for development.
	`,
	Aliases: []string{"profile"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("profiles called")
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
