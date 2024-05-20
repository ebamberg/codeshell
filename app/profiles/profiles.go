package profiles

import (
	"codeshell/applications"
	"codeshell/config"
	"codeshell/output"
	"codeshell/utils"
	"fmt"

	"github.com/spf13/viper"
)

type Profile struct {
	Id              string
	Displayname     string            `mapstructure:"displayName"`
	EnvVars         map[string]string `mapstructure:"envVars"`
	Applications    []string          `mapstructure:"applications"`
	AutoInstallApps bool              `mapstructure:"autoInstallApps"`
}

func getAllProfiles() (map[string]Profile, error) {
	var profiles = make(map[string]Profile)
	err := viper.UnmarshalKey(config.CONFIG_KEY_PROFILES, &profiles)
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
			//	log.Printf("set env variable %s = %s", strings.ToUpper(envVar), value)
			utils.SetEnvVariable(envVar, value)
		}

		ActivateApps(profile.Applications, profile.AutoInstallApps)
		return true
	} else {
		output.Errorf("profile [%s] not found", id)
		return false
	}
}

func ActivateApps(appList []string, autoInstallApps bool) {
	pi := output.NewProgressIndicator("activating applications...").Start()
	defer pi.Stop()

	//	installed := applications.ListInstalledAppications()
	for _, appKey := range appList {
		app, ok := applications.FindById(appKey)
		if ok && app.Status == applications.Installed {
			app.Activate()
		} else if ok && autoInstallApps && app.Status == applications.Available {
			err := applications.Install(app)
			app, ok := applications.FindById(appKey)
			if err == nil && ok {
				app.Activate()
			} else {
				output.Errorf("unable to install app %s. skipped activation..%s\n", err)
			}
		} else {
			output.Errorf("app %s is not installed. skipped activation.\n", appKey)
		}
	}

}
