package shell

import (
	"codeshell/config"
	"os/exec"
	"slices"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var cobraRootCmd *cobra.Command

func Run(rootCmd *cobra.Command) {
	cobraRootCmd = rootCmd
	printWelcomeMessage()
	// Create an interactive text input with single line input mode and show it

	inputPrompt := DefaultInteractivePromptInput
	// .WithOnInterruptFunc(func() {
	//	os.Exit(0)
	// })
	for {

		result, _ := inputPrompt.Show()
		if result == "quit" || result == "exit" {
			break
		}
		// Print a blank line for better readability
		pterm.Println()
		execute(result)
		// Print the user's answer with an info prefix
		//	pterm.Info.Printfln(": %s", result)
	}
}

func printWelcomeMessage() {

	clear()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithFullWidth().Println(config.GetString("terminal.style.title"))
	pterm.Println()
}

func clear() {
	print("\033[H\033[2J")
}

func execute(prompt string) {
	cmdArgs := strings.Split(prompt, " ")
	if isInternalCommand(cmdArgs) {
		cobraRootCmd.SetArgs(cmdArgs)
		err := cobraRootCmd.Execute()
		if err != nil {
			pterm.Error.Println(err)
		}
	} else {
		exe := cmdArgs[0]
		args := cmdArgs[1:]
		cmd := exec.Command(exe, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			pterm.Error.Println(err)
		} else {
			pterm.Println(string(out))
		}
	}
}

func isInternalCommand(args []string) bool {
	cmds := cobraRootCmd.Commands()
	return slices.ContainsFunc(cmds, func(cmd *cobra.Command) bool {
		if cmd.Use == args[0] {
			return true
		} else {
			return slices.Contains(cmd.Aliases, args[0])
		}
	})
}
