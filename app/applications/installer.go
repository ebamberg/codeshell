package applications

import (
	"codeshell/config"
	"codeshell/query"
	"codeshell/templating"
	"codeshell/vfs"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

func appPath(app Application) string {
	appsInstallpath := config.GetString(CONFIG_KEY_APP_PATH)
	return filepath.Join(appsInstallpath, app.Id, app.Version)
}

func UnInstall(newApp Application) error {

	if newApp.Status == Activated {
		return errors.New("Error uninstalling application. application is activated in current profile.")
	}
	appPath := appPath(newApp)
	err := os.RemoveAll(appPath)
	if err == nil {
		newApp.Status = Available
		localApps := localAppProvider.GetMapIndex()
		localApps[newApp.Id] = query.RemoveElement(localApps[newApp.Id], func(a Application) bool {
			return a.Id == newApp.Id && a.Version == newApp.Version && a.Status == Installed
		})
		config.Set(config.CONFIG_KEY_APPLICATIONS_INSTALLED, localApps)
	}

	return err
}

func Install(newApp Application) error {

	appPath := appPath(newApp)

	err := os.MkdirAll(appPath, 0)
	if err == nil {
		defer func() {
			if err != nil {
				os.RemoveAll(appPath)
			}
		}()
		// f, err := ioutil.TempFile("", "prefix")
		downloadFilePath := filepath.Join(appPath, "~download.zip")
		out, err := os.Create(downloadFilePath)
		if err == nil {
			defer os.Remove(downloadFilePath)
			defer out.Close()

			resp, err := http.Get(newApp.Source.Url)
			if err == nil {
				defer resp.Body.Close()
				var contentLength = 0
				clength := resp.Header.Get("Content-Length")
				if clength != "" {
					contentLength = int(contentLength)
				}
				_, err = vfs.Copy(out, resp.Body)
				//unzip file
				if err == nil {
					err = unzipSource(downloadFilePath, appPath, newApp.Source.IgnoreRootFolder, newApp.Source.archive)
					if err == nil {
						localApps := localAppProvider.GetMapIndex()

						installedApp := newApp // copy our struct
						installedApp.Path = appPath
						installedApp.BinaryPath = findBinaryPath(appPath)
						installedApp.Status = Installed
						// copy envVars from source to application
						installedApp.EnvVars = make(map[string]string)
						if newApp.Source.EnvVars != nil {
							for varName, varValue := range newApp.Source.EnvVars {
								installedApp.EnvVars[varName] = templating.ProcessPlaceholders(varValue, installedApp)
							}
						}

						localApps[installedApp.Id] = append(localApps[installedApp.Id], installedApp)
						config.Set(config.CONFIG_KEY_APPLICATIONS_INSTALLED, localApps)

					} else {
						return err
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
	return err
}
