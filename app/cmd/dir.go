/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/vfs"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var dirCmd = &cobra.Command{
	Use:     "dir",
	Aliases: []string{"ls"},
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

		if false == true {
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
		}

		output.PrintDirectoryTree(entries, func(row any) (int, string) {
			entry := row.(vfs.VFSEntry)
			var name string
			if entry.IsDir {
				name = pterm.Yellow(entry.Name)
			} else {
				name = entry.Name
			}
			level := len(strings.Split(entry.Path, string(os.PathSeparator))) - 1
			return level, name
		})

	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
}
