package applications

import (
	"codeshell/config"
	"codeshell/utils"
	"fmt"
	"os"
	"path/filepath"
)

const CONFIG_KEY_APP_PATH = "local.paths.applications"

type Status int

const (
	Available Status = iota + 1
	Installed
	Activated
	Remove
)

func (s Status) String() string {
	statuses := [...]string{"Available", "Installed", "Activated", "Remove"}
	if s < Available || s > Remove {
		return fmt.Sprintf("Status[%d]", int(s))
	}
	return statuses[s-1]
}

type Application struct {
	DisplayName string
	Path        string
	BinaryPath  string
	Status      Status
}

func ListInstalledAppications() map[string]Application {
	appspath := config.GetString(CONFIG_KEY_APP_PATH)
	entries, err := os.ReadDir(appspath)
	if err == nil {
		result := make(map[string]Application)
		for _, e := range entries {
			if e.IsDir() {
				fmt.Println(e.Name())
				app_name := e.Name()
				app_path := filepath.Join(appspath, app_name)
				bin_path := findBinaryPath(app_path)

				result[app_name] = Application{app_name, app_path, bin_path, Installed}
			}
		}
		return result
	} else {
		return nil
	}
}

func findBinaryPath(app_path string) string {
	bin_path := filepath.Join(app_path, "bin")
	if utils.DirectoryExists(bin_path) {
		return bin_path
	} else {
		return app_path
	}

}
