package profiles

import (
	"codeshell/applications"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const CONFIG_KEY_PROFILES = "profiles"

type Profile struct {
	Displayname  string            `mapstructure:"displayName"`
	EnvVars      map[string]string `mapstructure:"envVars"`
	Applications []string          `mapstructure:"applications"`
}

func getAllProfiles() (map[string]*Profile, error) {
	var profiles = make(map[string]*Profile)
	viper.UnmarshalKey(CONFIG_KEY_PROFILES, &profiles)
	return profiles, nil
}

func GetProfile(id string) (*Profile, error) {
	profiles, err := getAllProfiles()
	if err == nil {
		return profiles[id], nil
	} else {
		return nil, err
	}
}

/*
Activates a profile.

This reset the Path in env.
Sets all envvars from profile config.
prepends all applications mentioned in profile to the path
*/
func ActivateProfile(id string) error {
	profile, err := GetProfile(id)
	if err == nil {
		if profile != nil {
			resetEnvPath()
			for envVar, value := range profile.EnvVars {
				log.Printf("set env variable %s = %s", strings.ToUpper(envVar), value)
				setEnvVariable(envVar, value)
			}
			ActivateApps(profile.Applications)
			return nil
		} else {
			return fmt.Errorf("profile [%s] not found", id)
		}
	} else {
		return err
	}
}

/*
sets an env variable and take care of special cases
PATH variable is not overridden but get prepended as suffix to the  Path
when setting the PATH for the first time the original state is stored in
a env var. resetEnvPath can be used afterward to reset to that state
*/
func setEnvVariable(envVar string, value string) {
	envVar = strings.ToUpper(envVar)
	if envVar == "PATH" {
		current := os.Getenv("PATH")
		originalPath := os.Getenv("CODESHELL_ORIGINAL_PATH")
		if originalPath == "" {
			os.Setenv("CODESHELL_ORIGINAL_PATH", current)
		}
		value = value + string(os.PathListSeparator) + current
	}
	os.Setenv(envVar, value)
}

func appendEnvPath(path string) {
	setEnvVariable("PATH", path)
}

/*
resets the envvariable PATH to the original state
before it has been modified by CODESHELL
*/
func resetEnvPath() {
	originalPath := os.Getenv("CODESHELL_ORIGINAL_PATH")
	if originalPath != "" {
		os.Setenv("PATH", originalPath)
	}
}

func ActivateApps(appList []string) {
	fmt.Println("activating applications...")

	installed := applications.ListInstalledAppications()
	for _, appKey := range appList {
		if app, ok := installed[appKey]; ok {

			appendEnvPath(app.BinaryPath)
			fmt.Printf("activated : \t%s\t\t%s\n", app.DisplayName, app.Path)

		} else {
			fmt.Printf("app %s is not installed. skipped activation.\n", appKey)
		}
	}

}
