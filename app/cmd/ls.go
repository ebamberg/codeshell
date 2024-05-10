/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/vfs"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"dir"},
	Short:   "List content of directory",
	Long: `List the content of directory. 
	If no path is given then this command returns the content of the current working directory.
	Examples:
		ls               -	lists the content of current working directory
		ls /home/users/  - list the content of the directory /home/users/`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) > 0 {
			path = args[0]
		} else {
			path = "./"
		}
		entries := vfs.DefaultFilesystem.List(path)

		output.PrintTidySlice(entries, []string{"Name"}, func(row any) []string {
			entry := row.(vfs.VFSEntry)
			var name string
			if entry.IsDir {
				name = pterm.Yellow(entry.Name)
			} else {
				name = entry.Name
			}
			return []string{name}
		})
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
