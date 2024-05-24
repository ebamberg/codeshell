package applications

import (
	"codeshell/config"
	"codeshell/output"
	"os"
	"path/filepath"
	"slices"
)

type LocalOnFilesystemInstalledApplicationProvider struct {
}

func (this *LocalOnFilesystemInstalledApplicationProvider) GetMapIndex() map[string][]Application {
	appspath := config.GetString(CONFIG_KEY_APP_PATH)
	entries, err := os.ReadDir(appspath)
	if err == nil {
		activated := getActivatedAppIds()
		result := make(map[string][]Application)
		for _, e := range entries {
			if e.IsDir() {
				id := e.Name()
				app_name := e.Name()
				app_path := filepath.Join(appspath, app_name)
				// list version
				versions, err := os.ReadDir(app_path)
				if err == nil {
					for _, v := range versions {
						if v.IsDir() {
							version := v.Name()
							id = id
							version_path := filepath.Join(app_path, version)
							bin_path := findBinaryPath(version_path)
							var status Status
							if slices.Contains(activated, id) {
								//TODO check for version as well
								status = Activated
							} else {
								status = Installed
							}

							result[id] = append(result[id], Application{Id: id, DisplayName: app_name, Version: version, Path: app_path, BinaryPath: bin_path, Status: status})
						}
					}
				} else {
					output.Errorln(err)
				}
			}
		}
		return result
	} else {
		return nil
	}

}

func (this *LocalOnFilesystemInstalledApplicationProvider) List() []Application {
	return FlattenMap(this.GetMapIndex())
}
