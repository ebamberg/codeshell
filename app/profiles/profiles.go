package profiles

import (
	"codeshell/applications"
	"codeshell/config"
	"codeshell/events"
	"codeshell/output"
	"codeshell/utils"
	"codeshell/vfs"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Id              string
	Displayname     string            `mapstructure:"displayName"`
	CheckVirtualEnv bool              `mapstructure:"checkVirtualEnv"`
	EnvVars         map[string]string `mapstructure:"envVars"`
	Applications    []string          `mapstructure:"applications"`
	AutoInstallApps bool              `mapstructure:"autoInstallApps"`
}

var CurrentProfile *Profile

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
		setEnvVariables(profile.EnvVars)

		ActivateApps(profile.Applications, profile.AutoInstallApps)
		CurrentProfile = &profile
		events.Broadcast(events.ApplicationEvent{Eventtype: events.PROFILE_ACTIVATED, Payload: profile})
		return true
	} else {
		output.Errorf("profile [%s] not found", id)
		return false
	}
}

func Import(id string) error {
	imp := make(map[string]Profile, 0)
	repo, err := findRepo()
	if err != nil {
		return err
	}
	if repo != nil {
		file, err := repo.Read("profiles.yaml")
		if err == nil {
			defer file.Close()
			buf, err := ioutil.ReadAll(file)

			if err == nil {
				err = yaml.Unmarshal(buf, &imp)
				if err == nil {
					profiles, err := getAllProfiles()
					if err == nil {
						if _, exists := profiles[id]; !exists {

							if newProfile, exists := imp[id]; exists {
								profiles[id] = newProfile
								config.Set(config.CONFIG_KEY_PROFILES, profiles)
								return nil
							} else {
								return errors.New("profile not found in repository")
							}
						} else {
							return errors.New("profile already exits locally")
						}
					} else {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		}
	} else {
		err = errors.New("no repository configured")
	}
	return err

}

func findRepo() (vfs.VFS, error) {
	repo_url := config.GetString(config.CONFIG_KEY_REPO_APP_URL)
	if repo_url != "" {
		return vfs.FromUrlString(repo_url)
	} else {
		return nil, nil
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

func setEnvVariables(envVars map[string]string) {
	for envVar, value := range envVars {
		//	log.Printf("set env variable %s = %s", strings.ToUpper(envVar), value)
		utils.SetEnvVariable(envVar, value)
	}
}
