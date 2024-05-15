package applications

import (
	"codeshell/config"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var available = map[string]Application{
	"eclipse-java": Application{Id: "eclipse-java", DisplayName: "Eclipse for Java and WebDev", Status: Available, Version: "2024-03", source: appInstallationSource{size: 0, url: "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-jee-2024-03-R-win32-x86_64.zip"}},
	"eclipse-cpp":  Application{Id: "eclipse-cpp", DisplayName: "Eclipse for C/C++ developer", Status: Available, Version: "2024-03", source: appInstallationSource{size: 0, url: "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-cpp-2024-03-R-win32-x86_64.zip"}},
	"maven":        Application{Id: "maven", DisplayName: "Apache Maven", Status: Available, Version: "3.9.6", source: appInstallationSource{size: 9513253, url: "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.zip"}},
	"npp":          Application{Id: "npp", DisplayName: "Notepad++", Status: Available, Version: "8.6.7", source: appInstallationSource{size: 5998909, url: "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.7/npp.8.6.7.portable.x64.zip"}},
}

func ListAvailableApplications() map[string]Application {
	return available
}

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

			resp, err := http.Get(newApp.source.url)
			if err == nil {
				defer resp.Body.Close()
				_, err = io.Copy(out, resp.Body)
				//unzip file
				if err == nil {
					err = unzipSource(downloadFilePath, appPath)
				}
			}
		}
	}
	return err
}
