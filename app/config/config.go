package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create a new default config file
			viper.WriteConfigAs("./codeshell.config")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

}
