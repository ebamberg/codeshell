/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"strings"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "print a text to the console.",
	Long:  `print a text to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		output.Println(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
