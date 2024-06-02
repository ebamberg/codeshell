/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/utils"
	"strings"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"export"},
	Short:   "set the value of an environment variable",
	Long: `set the value of an environment variable. For example:

	set myvar=helloworld
	set myvar helloworld
	.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		var varName, varValue string
		var valid bool
		if len(args) == 1 {
			varName, varValue, valid = strings.Cut(args[0], "=")

		} else {
			varName = args[0]
			varValue = args[1]
			valid = true
		}
		if valid {
			utils.SetEnvVariable(varName, varValue)
			output.Infof("variabe %s set", varName)
		} else {
			output.Errorln("usage 'set varName=varValue'")
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

}
