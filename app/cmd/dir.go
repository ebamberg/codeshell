/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/style"
	"codeshell/vfs"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var depth *int

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
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		//	depth, _ := cmd.Flags().GetInt("depth")
		var path string
		if len(args) > 0 {
			path = args[0]
		} else {
			path = "./"
		}
		entries := vfs.DefaultFilesystem.List(path, *depth)

		if false == true {
			output.PrintTidySlice(entries, []string{"Name"}, func(row any) []string {
				entry := row.(vfs.VFSEntry)
				var name string
				if entry.IsDir {
					name = style.Folder(entry.Name)
				} else {
					name = entry.Name
				}
				return []string{name}
			})
		}

		vfs.PrintDirectoryTree(entries, func(e vfs.VFSEntry) (int, string) {
			var name string
			if e.IsDir {
				name = pterm.Yellow(e.Name)
			} else {
				name = e.Name
			}
			level := len(strings.Split(e.Path, string(os.PathSeparator))) - 1
			return level, name
		})

	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
	depth = dirCmd.Flags().Int("depth", 0, "recursive showing directory to depth")
}
