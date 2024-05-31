package style

import (
	"codeshell/applications"

	"github.com/pterm/pterm"
)

func Folder(a string) string {
	return pterm.Yellow(a)
}

func Prompt(prompt string) string {
	return pterm.Red(prompt)
}

func Profile(displayname string) string {
	return pterm.BgLightRed.Sprint(displayname)
}

func AppStatus(s applications.Status) string {
	switch s {
	case applications.Available:
		return pterm.LightBlue(s.String())
	case applications.Installed:
		return pterm.Yellow(s.String())
	case applications.Activated:
		return pterm.Green(s.String())
	default:
		return s.String()
	}

}
