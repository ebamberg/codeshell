/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/events"
	"codeshell/output"
	"codeshell/vfs"

	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "change the current working directory.",
	Long:  `change the current working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			path := args[0]
			err := vfs.DefaultFilesystem.Chdir(path)
			if err != nil {
				output.Errorln(err)
			}
			events.Broadcast(events.ApplicationEvent{events.FS_WORKDIR_CHANGED, path})
		} else {
			output.Errorln("please pass path the new working directory.")
			output.Infoln(cmd.UsageString())
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
