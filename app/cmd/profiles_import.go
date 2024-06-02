/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"codeshell/output"
	"codeshell/profiles"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var profilesImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a profile from the remoterepository",
	Long: `import a profile from the remote repository and add it the list of local profiles.
	example:
		profile import java17
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := profiles.Import(args[0])
		if err == nil {
			output.Successf("profile %s imported.", args[0])
		} else {
			output.Errorln(err)
		}
	},
}

func init() {
	profilesCmd.AddCommand(profilesImportCmd)
}
