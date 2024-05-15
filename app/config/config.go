package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const CONFIG_KEY_APP_PATH = "local.paths.applications"

var defaults = map[string]string{
	"local.paths.applications": "./apps/",
	"terminal.style.title":     "Codeshell",
}

func setDefaults() {
	for key, val := range defaults {
		viper.SetDefault(key, val)
	}
}

// initConfig reads in config file and ENV variables if set.
func Init(configLocations ...string) {

	if len(configLocations) > 0 {
		// Use config file from the flag.
		for _, cfgFile := range configLocations {
			viper.SetConfigFile(cfgFile)
		}
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in working and home directory with name "codeshell.config" (without extension).

		viper.AddConfigPath("./")
		viper.AddConfigPath(home)

		viper.SetConfigType("yaml")
		viper.SetConfigName("codeshell.config")
		viper.SetEnvPrefix("CODESHELL_")
	}

	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create a new default config file
			viper.WriteConfigAs("./codeshell.config")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	createAppPath()

}

func createAppPath() {

	appspath := GetString(CONFIG_KEY_APP_PATH)
	fmt.Println(">>" + appspath)
	err := os.MkdirAll(appspath, 0755)
	if err != nil {
		panic(err)
	}
}

func GetString(param string) string {
	return viper.GetString(param)
}

func Set(path string, value any) {
	viper.Set(path, value)
}
