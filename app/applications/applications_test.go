package applications

import (
	"codeshell/config"
	"codeshell/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestAppFolder() string {
	appsdir, _ := os.MkdirTemp("", "apps")
	os.MkdirAll(filepath.Join(appsdir, "test1", "0.0.1", "bin"), os.FileMode(0777))
	os.MkdirAll(filepath.Join(appsdir, "test2", "0.0.1"), os.FileMode(0777))

	config.Set(CONFIG_KEY_APP_PATH, appsdir)
	return appsdir
}

func teardownTestAppFolder(appsdir string) {
	os.RemoveAll(appsdir)
}

func TestListInstalledApplications(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	applicationsfound := ListInstalledAppications()

	assert.Equal(t, 2, len(applicationsfound))

}

func TestListInstalledAppications_have_correct_bin_folder(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	applicationsfound := ListInstalledAppications()

	assert.True(t, strings.HasSuffix(applicationsfound["test1"][0].BinaryPath, filepath.Join("0.0.1", "bin")))
	assert.True(t, strings.HasSuffix(applicationsfound["test2"][0].BinaryPath, filepath.Join("test2", "0.0.1")))
}

func TestActivateApp(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	applicationsfound := ListInstalledAppications()
	applicationsfound["test1"][0].Activate()

	assert.Contains(t, strings.Split(utils.GetEnvVariable(ENV_KEY_ACTIVACTED), ","), "test1")
}

func TestActivateApp_twice_doesnt_add_app_twice(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	applicationsfound := ListInstalledAppications()
	applicationsfound["test1"][0].Activate()
	applicationsfound["test1"][0].Activate()
	ac := strings.Split(utils.GetEnvVariable(ENV_KEY_ACTIVACTED), ",")
	var count int
	for _, a := range ac {
		if a == "test1" {
			count++
		}
	}
	assert.Equal(t, 1, count)
}

func TestListInstalledApplications_have_correct_status(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	origActivated := utils.GetEnvVariable(ENV_KEY_ACTIVACTED)
	utils.SetEnvVariable(ENV_KEY_ACTIVACTED, "test2,test3")
	defer utils.SetEnvVariable(ENV_KEY_ACTIVACTED, origActivated)
	applicationsfound := ListInstalledAppications()

	assert.Equal(t, Installed, applicationsfound["test1"][0].Status)
	assert.Equal(t, Activated, applicationsfound["test2"][0].Status)
}

func TestApplicationsStatus(t *testing.T) {
	assert.Equal(t, "Available", Available.String())
	assert.Equal(t, "Installed", Installed.String())
	assert.Equal(t, "Activated", Activated.String())
	assert.Equal(t, "Remove", Remove.String())
	assert.Equal(t, "Status[100]", Status(100).String())
}
