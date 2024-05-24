package applications

import (
	"codeshell/ext"
	"codeshell/output"
	"codeshell/query"
	"codeshell/utils"
	"fmt"
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

type ApplicationProvider interface {
	//	ext.Provider[Application]
	GetMapIndex() map[string][]Application
	List() []Application
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
	return ListApplicationsFilteredBy(ext.MatchAll)
}

func ListApplicationsFilteredBy(matches ext.Predicate[Application]) map[string][]Application {
	apps := make(map[string][]Application, 0)
	available := Providers[0].GetMapIndex()
	installed := Providers[1].GetMapIndex()

	// maps.Copy(apps, available)
	// make a deep copy that also create new slices
	for k, sli := range available {
		tmp := make([]Application, 0, len(sli))
		for _, a := range sli {
			if matches(a) {
				apps[k] = append(tmp, sli...)
			}
		}
	}

	//	maps.Copy(apps, installed)
	for key, val := range installed {
		if versions, isMapContainsKey := apps[key]; isMapContainsKey {
			for _, inst := range val {
				if matches(inst) {
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
			}
		} else {
			apps[key] = val
		}
	}
	return apps
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
