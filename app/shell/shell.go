package shell

import (
	"bufio"
	"codeshell/config"
	"codeshell/output"
	"codeshell/profiles"
	"codeshell/style"
	"codeshell/vfs"
	"io"
	"os"
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
		err := Execute(result)
		if err != nil {
			output.Errorln(err)
		}
	}
}

func prompt() string {
	var prompt []string
	workdir, _ := vfs.DefaultFilesystem.Getwd()
	profile := profiles.CurrentProfile
	if profile != nil {
		prompt = append(prompt, style.Profile(" "+profile.Displayname+" "))
	}
	venv := os.Getenv("VIRTUAL_ENV")
	if venv != "" {
		prompt = append(prompt, style.PythonVenv(" "+venv+" "))
	}

	prompt = append(prompt, style.Prompt(workdir))
	return strings.Join(prompt, " ")
}

func printWelcomeMessage() {

	clear()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightRed)).WithFullWidth().Println(config.GetString("terminal.style.title"))
	pterm.Println()
}

func clear() {
	print("\033[H\033[2J")
}

func ExecuteScript(path string) error {
	reader, err := vfs.DefaultFilesystem.Read(path)
	if err == nil {
		defer reader.Close()
		return ExecuteBatch(reader)
	} else {
		return err
	}

}

func ExecuteBatch(reader io.Reader) error {
	fileScanner := bufio.NewScanner(reader)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		err := Execute(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(prompt string) error {
	cmdArgs := strings.Split(prompt, " ")
	if cmdArgs[0] == "help" {
		output.Infoln(cobraRootCmd.UsageString())
	} else if isInternalCommand(cmdArgs) {
		cobraRootCmd.SetArgs(cmdArgs)
		defer resetSubCommandFlagValues(cobraRootCmd)
		//	inbuf := bytes.NewBufferString("")
		//	errbuf := bytes.NewBufferString("")
		//	cobraRootCmd.SetOut(inbuf)
		//	cobraRootCmd.SetErr(errbuf)s
		err := cobraRootCmd.Execute()
		if err != nil {
			return err
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
		output.Println("")
	} else {
		exe := cmdArgs[0]
		args := cmdArgs[1:]
		cmd := exec.Command(exe, args...)

		out, err := cmd.CombinedOutput()
		if err != nil {
			output.Println(string(out))
			return err
		} else {
			output.Println(string(out))
		}
	}
	return nil
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
