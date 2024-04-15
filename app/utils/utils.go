package utils

import "os"

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
