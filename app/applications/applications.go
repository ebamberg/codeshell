package applications

import (
	"codeshell/config"
	"codeshell/output"
	"codeshell/query"
	"codeshell/utils"
	"fmt"
	"maps"
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

type appInstallationSource struct {
	url              string
	size             int
	ignoreRootFolder bool
}

type Application struct {
	Id          string
	DisplayName string
	Path        string
	BinaryPath  string
	Status      Status
	Version     string
	source      appInstallationSource
}

func (this Application) Activate() {
	utils.AppendEnvPath(this.BinaryPath)
	activated := strings.Split(utils.GetEnvVariable(ENV_KEY_ACTIVACTED), ",")
	if !slices.Contains(activated, this.Id) {
		activated = append(activated, this.Id)
		utils.SetEnvVariable(ENV_KEY_ACTIVACTED, strings.Join(activated, ","))
	}
	output.Infof("activated : \t%s\t\t%s\n", this.DisplayName, this.Path)

}

func getActivatedAppIds() []string {
	activated := utils.GetEnvVariable(ENV_KEY_ACTIVACTED)
	return strings.Split(activated, ",")
}

func FindById(identifier string) (Application, bool) {
	idElements := strings.Split(identifier, ":")
	var id string
	var version string
	id = idElements[0]
	if len(idElements) > 1 {
		version = idElements[1]
	}
	allApps := ListApplications()
	if newApps, isMapContainsKey := allApps[id]; isMapContainsKey {
		if version == "" {
			return newApps[0], true
		} else {
			byVersion := query.Filter(newApps, func(a Application) bool {
				return (a.Version == version || version == "")
			})
			if len(byVersion) > 0 {
				return byVersion[0], true
			} else {
				return Application{}, false
			}
		}

	} else {
		return Application{}, false
	}
}

func ListApplications() map[string][]Application {
	apps := make(map[string][]Application, 0)
	available := ListAvailableApplications()
	installed := ListInstalledAppications()

	maps.Copy(apps, available)
	maps.Copy(apps, installed)
	return apps
}

func ListInstalledAppications() map[string][]Application {
	appspath := config.GetString(CONFIG_KEY_APP_PATH)
	entries, err := os.ReadDir(appspath)
	if err == nil {
		activated := getActivatedAppIds()
		result := make(map[string][]Application)
		for _, e := range entries {
			if e.IsDir() {
				id := e.Name()
				app_name := e.Name()
				app_path := filepath.Join(appspath, app_name)
				// list version
				versions, err := os.ReadDir(app_path)
				if err == nil {
					for _, v := range versions {
						if v.IsDir() {
							version := v.Name()
							id = id
							version_path := filepath.Join(app_path, version)
							bin_path := findBinaryPath(version_path)
							var status Status
							if slices.Contains(activated, id) {
								//TODO check for version as well
								status = Activated
							} else {
								status = Installed
							}

							result[id] = append(result[id], Application{Id: id, DisplayName: app_name, Version: version, Path: app_path, BinaryPath: bin_path, Status: status})
						}
					}
				} else {
					output.Errorln(err)
				}
			}
		}
		return result
	} else {
		return nil
	}
}

func FlattenMap(appMap map[string][]Application) []Application {
	v := make([]Application, 0, len(appMap))

	for _, values := range appMap {
		for _, value := range values {
			v = append(v, value)
		}
	}
	return v
}

func findBinaryPath(app_path string) string {
	bin_path := filepath.Join(app_path, "bin")
	if utils.DirectoryExists(bin_path) {
		return bin_path
	} else {
		return app_path
	}

}
