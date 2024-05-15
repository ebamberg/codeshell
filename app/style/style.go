package style

import (
	"codeshell/applications"

	"github.com/pterm/pterm"
)

func Folder(a string) string {
	return pterm.Yellow(a)
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
