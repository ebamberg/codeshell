package applications

import (
	"codeshell/config"
	"codeshell/utils"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const CONFIG_KEY_APP_PATH = "local.paths.applications"
const ENV_KEY_ACTIVACTED = "CS_APP_ACTIVATED"

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

type application struct {
	id          string
	DisplayName string
	Path        string
	BinaryPath  string
	Status      Status
}

func (this application) Activate() {
	utils.AppendEnvPath(this.BinaryPath)
	activated := strings.Split(utils.GetEnvVariable(ENV_KEY_ACTIVACTED), ",")
	if !slices.Contains(activated, this.id) {
		activated = append(activated, this.id)
		utils.SetEnvVariable(ENV_KEY_ACTIVACTED, strings.Join(activated, ","))
	}
	fmt.Printf("activated : \t%s\t\t%s\n", this.DisplayName, this.Path)

}

func getActivatedAppIds() []string {
	activated := utils.GetEnvVariable(ENV_KEY_ACTIVACTED)
	return strings.Split(activated, ",")
}

func ListInstalledAppications() map[string]application {
	appspath := config.GetString(CONFIG_KEY_APP_PATH)
	entries, err := os.ReadDir(appspath)
	if err == nil {
		activated := getActivatedAppIds()
		result := make(map[string]application)
		for _, e := range entries {
			if e.IsDir() {
				fmt.Println(e.Name())
				id := e.Name()
				app_name := e.Name()
				app_path := filepath.Join(appspath, app_name)
				bin_path := findBinaryPath(app_path)
				var status Status
				if slices.Contains(activated, app_name) {
					status = Activated
				} else {
					status = Installed
				}
				result[app_name] = application{id, app_name, app_path, bin_path, status}
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
