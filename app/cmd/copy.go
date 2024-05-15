/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/vfs"
	"io"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{"cp"},
	Short:   "copy content from a location to another",
	Long:    `copy content from a location to another.`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			src := args[0]
			dest := args[1]
			var err error
			in, err := vfs.DefaultFilesystem.Read(src)
			if err == nil {
				defer in.Close()
				out, err := vfs.DefaultFilesystem.Create(dest)
				if err == nil {
					defer out.Close()
					_, err := io.Copy(out, in)
					if err == nil {
						output.Infof("content copied from %s to %s\n", src, dest)
					}
				}
			}
			if err != nil {
				output.Errorln(err)
			}

		} else {
			output.Errorln("no source or destination location defined.")
			output.Infoln(cmd.UsageString())
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
