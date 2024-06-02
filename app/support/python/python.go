package python

import (
	"codeshell/events"
	"codeshell/output"
	"codeshell/profiles"
	"codeshell/utils"
	"codeshell/vfs"
	"path/filepath"
)

var OLD_VIRTUAL_ENV string

func ActivateVirtualEnvironment(path string) {
	if OLD_VIRTUAL_ENV != "" {
		DeactivateVirtualEnvironment()
	}
	OLD_VIRTUAL_ENV = utils.GetEnvVariable("VIRTUAL_ENV")
	utils.SetEnvVariable("VIRTUAL_ENV", path)
	utils.AppendEnvPath(filepath.Join(path, "Scripts"))
}

func DeactivateVirtualEnvironment() {
	path := utils.GetEnvVariable("VIRTUAL_ENV")
	utils.SetEnvVariable("VIRTUAL_ENV", OLD_VIRTUAL_ENV)
	utils.RemoveEnvPath(filepath.Join(path, "Scripts"))
}

func init() {
	events.RegisterListener(func(e events.ApplicationEvent) {
		if e.Eventtype == events.FS_WORKDIR_CHANGED {
			if profiles.CurrentProfile != nil && profiles.CurrentProfile.CheckVirtualEnv == true {
				if vfs.DefaultFilesystem.Exists("venv") {
					workDir, err := vfs.DefaultFilesystem.Getwd()
					if err == nil {
						ActivateVirtualEnvironment(filepath.Join(workDir, "venv"))
						output.Infoln("Virtual Python Environment activated")
					} else {
						output.Errorln(err)
					}
				}
			}
		}
	})
}
