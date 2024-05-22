package applications

import (
	"codeshell/config"
	"codeshell/output"
	"codeshell/query"
	"codeshell/utils"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/viper"
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
	envVars          map[string]string
}

type Application struct {
	Id          string
	DisplayName string
	Path        string
	BinaryPath  string
	Status      Status
	Version     string
	EnvVars     map[string]string `mapstructure:"envVars"`
	source      appInstallationSource
}

func (this Application) Activate() {
	utils.AppendEnvPath(this.BinaryPath)
	activated := strings.Split(utils.GetEnvVariable(ENV_KEY_ACTIVACTED), ",")
	if !slices.Contains(activated, this.Id) {
		activated = append(activated, this.Id)
		utils.SetEnvVariable(ENV_KEY_ACTIVACTED, strings.Join(activated, ","))
	}
	setEnvVariables(this.EnvVars)
	output.Infof("activated : \t%s\t\t%s\n", this.DisplayName, this.Path)

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
	if newApps, isMapContainsKey := allApps[id]; isMapContainsKey && len(newApps) > 0 {
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

	// maps.Copy(apps, available)
	// make a deep copy that also create new slices
	for k, sli := range available {
		tmp := make([]Application, 0, len(sli))
		apps[k] = append(tmp, sli...)
	}

	//	maps.Copy(apps, installed)
	for key, val := range installed {
		if versions, isMapContainsKey := apps[key]; isMapContainsKey {
			for _, inst := range val {
				j := slices.IndexFunc(versions, func(a Application) bool {
					return a.Version == inst.Version
				})
				if j > -1 {
					// override available entry with installed
					versions[j] = inst
				} else {
					// just appendinstalled version // this should never happen
					apps[key] = append(apps[key], inst)
				}
			}
		} else {
			apps[key] = val
		}
	}
	return apps
}

func getLocalApplications() (map[string][]Application, error) {
	apps := make(map[string][]Application, 0)
	err := viper.UnmarshalKey(config.CONFIG_KEY_APPLICATIONS_INSTALLED, &apps)
	return apps, err
}

func ListInstalledAppications() map[string][]Application {
	apps, err := getLocalApplications()
	if err == nil {
		return apps
	} else {
		panic(err)
	}
	// appspath := config.GetString(CONFIG_KEY_APP_PATH)
	// entries, err := os.ReadDir(appspath)
	// if err == nil {
	// 	activated := getActivatedAppIds()
	// 	result := make(map[string][]Application)
	// 	for _, e := range entries {
	// 		if e.IsDir() {
	// 			id := e.Name()
	// 			app_name := e.Name()
	// 			app_path := filepath.Join(appspath, app_name)
	// 			// list version
	// 			versions, err := os.ReadDir(app_path)
	// 			if err == nil {
	// 				for _, v := range versions {
	// 					if v.IsDir() {
	// 						version := v.Name()
	// 						id = id
	// 						version_path := filepath.Join(app_path, version)
	// 						bin_path := findBinaryPath(version_path)
	// 						var status Status
	// 						if slices.Contains(activated, id) {
	// 							//TODO check for version as well
	// 							status = Activated
	// 						} else {
	// 							status = Installed
	// 						}

	// 						result[id] = append(result[id], Application{Id: id, DisplayName: app_name, Version: version, Path: app_path, BinaryPath: bin_path, Status: status})
	// 					}
	// 				}
	// 			} else {
	// 				output.Errorln(err)
	// 			}
	// 		}
	// 	}
	// 	return result
	// } else {
	// 	return nil
	// }
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

func setEnvVariables(envVars map[string]string) {
	for envVar, value := range envVars {
		utils.SetEnvVariable(envVar, value)
	}
}

func getActivatedAppIds() []string {
	activated := utils.GetEnvVariable(ENV_KEY_ACTIVACTED)
	return strings.Split(activated, ",")
}
