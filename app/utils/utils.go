package utils

import (
	"os"
	"strings"
)

/*
check if a directory exists
return false if the directory doen't exists or the filepath points to a file instead of a directory
*/
func DirectoryExists(path string) bool {

	if len(path) == 0 {
		return false
	}
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) || !fileinfo.IsDir() {
		return false
	} else {
		return true
	}
}

/*
sets an env variable and take care of special cases
PATH variable is not overridden but get prepended as suffix to the  Path
when setting the PATH for the first time the original state is stored in
a env var. resetEnvPath can be used afterward to reset to that state
*/
func SetEnvVariable(envVar string, value string) {
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

func AppendEnvPath(path string) {
	SetEnvVariable("PATH", path)
}

/*
resets the envvariable PATH to the original state
before it has been modified by CODESHELL
*/
func ResetEnvPath() {
	originalPath := os.Getenv("CODESHELL_ORIGINAL_PATH")
	if originalPath != "" {
		os.Setenv("PATH", originalPath)
	}
}
