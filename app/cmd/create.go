/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:        "create [flags] [profile_name]",
	Short:      "creates a new empty profile and activates it.",
	Long:       `creates a new empty profile and activates it.`,
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"profile_name"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called", args)
	},
}

func init() {
	profilesCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//	createCmd.PersistentFlags().String("profileName", "empty", "the name of the profile to create")
}
