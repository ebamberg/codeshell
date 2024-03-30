package profiles

import (
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
	Applications string            `mapstructure:"applications"`
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

func ActivateProfile(id string) error {
	profile, err := GetProfile(id)
	if err == nil {
		if profile != nil {
			for envVar, value := range profile.EnvVars {
				log.Printf("set env variable %s = %s", strings.ToUpper(envVar), value)
				os.Setenv(strings.ToUpper(envVar), value)
			}
			return nil
		} else {
			return fmt.Errorf("profile [%s] not found", id)
		}
	} else {
		return err
	}
}
