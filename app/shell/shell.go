package shell

import (
	"codeshell/config"
	"codeshell/output"
	"codeshell/profiles"
	"codeshell/style"
	"codeshell/vfs"
	"os/exec"
	"slices"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var cobraRootCmd *cobra.Command

func Run(rootCmd *cobra.Command) {
	cobraRootCmd = rootCmd
	printWelcomeMessage()
	// Create an interactive text input with single line input mode and show it

	inputPrompt := DefaultInteractivePromptInput.WithTextStyle(pterm.NewStyle(pterm.FgRed))
	// .WithOnInterruptFunc(func() {
	//	os.Exit(0)
	// })
	for {
		result, _ := inputPrompt.Show(prompt())
		if result == "quit" || result == "exit" {
			break
		}
		// Print a blank line for better readability
		output.Println("")
		execute(result)
		// Print the user's answer with an info prefix
		//	pterm.Info.Printfln(": %s", result)
	}
}

func prompt() string {
	var prompt []string
	workdir, _ := vfs.DefaultFilesystem.Getwd()
	profile := profiles.CurrentProfile
	if profile != nil {
		prompt = append(prompt, style.Profile(" "+profile.Displayname+" "))
	}
	prompt = append(prompt, style.Prompt(workdir))
	return strings.Join(prompt, " ")
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
	if cmdArgs[0] == "help" {
		output.Infoln(cobraRootCmd.UsageString())
	} else if isInternalCommand(cmdArgs) {
		cobraRootCmd.SetArgs(cmdArgs)
		//	inbuf := bytes.NewBufferString("")
		//	errbuf := bytes.NewBufferString("")
		//	cobraRootCmd.SetOut(inbuf)
		//	cobraRootCmd.SetErr(errbuf)s
		err := cobraRootCmd.Execute()
		if err != nil {
			output.Errorln(err)
			//		} else {
			//			sIn, e1 := io.ReadAll(inbuf)
			//			sErr, e2 := io.ReadAll(errbuf)
			//			if len(sIn) > 0 {
			//				fmt.Println(sIn)
			//			}
			//			if len(sErr) > 0 {
			//				output.Errorln(string(sErr))
			//			}
			//			if e1 != nil {
			//				output.Errorln(e1)
			//			}
			//			if e1 != nil {
			//				output.Errorln(e2)
			//			}

		}
		resetSubCommandFlagValues(cobraRootCmd)
		output.Println("")
	} else {
		exe := cmdArgs[0]
		args := cmdArgs[1:]
		cmd := exec.Command(exe, args...)

		out, err := cmd.CombinedOutput()
		if err != nil {
			output.Errorln(err)
		} else {
			output.Println(string(out))
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

func resetSubCommandFlagValues(root *cobra.Command) {
	for _, c := range root.Commands() {
		c.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				f.Value.Set(f.DefValue)
				f.Changed = false
			}
		})
		resetSubCommandFlagValues(c)
	}
}
