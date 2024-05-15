package profiles

import (
	"codeshell/applications"
	"codeshell/output"
	"codeshell/utils"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const CONFIG_KEY_PROFILES = "profiles"

type Profile struct {
	Displayname  string            `mapstructure:"displayName"`
	EnvVars      map[string]string `mapstructure:"envVars"`
	Applications []string          `mapstructure:"applications"`
}

func getAllProfiles() (map[string]Profile, error) {
	var profiles = make(map[string]Profile)
	err := viper.UnmarshalKey(CONFIG_KEY_PROFILES, &profiles)
	return profiles, err
}

func ListProfiles() map[string]Profile {
	profiles, err := getAllProfiles()
	if err == nil {
		return profiles
	} else {
		return make(map[string]Profile, 0)
	}
}

func GetProfile(id string) (Profile, bool) {
	profiles, err := getAllProfiles()
	if err != nil {
		fmt.Println(err)
	}
	profile, exists := profiles[id]
	return profile, exists

}

/*
Activates a profile.

This reset the Path in env.
Sets all envvars from profile config.
prepends all applications mentioned in profile to the path
*/
func ActivateProfile(id string) bool {
	profile, exists := GetProfile(id)
	if exists {
		utils.ResetEnvPath()
		for envVar, value := range profile.EnvVars {
			log.Printf("set env variable %s = %s", strings.ToUpper(envVar), value)
			utils.SetEnvVariable(envVar, value)
		}
		ActivateApps(profile.Applications)
		return true
	} else {
		output.Errorf("profile [%s] not found", id)
		return false
	}
}

func ActivateApps(appList []string) {
	output.Println("activating applications...")

	installed := applications.ListInstalledAppications()
	for _, appKey := range appList {
		if app, ok := installed[appKey]; ok {
			app.Activate()
		} else {
			fmt.Printf("app %s is not installed. skipped activation.\n", appKey)
		}
	}

}
