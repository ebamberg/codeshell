package applications

import (
	"codeshell/config"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestAppFolder() string {
	appsdir, _ := os.MkdirTemp("", "apps")
	os.Mkdir(filepath.Join(appsdir, "test1"), os.FileMode(0777))
	os.Mkdir(filepath.Join(appsdir, "test1", "bin"), os.FileMode(0777))
	os.Mkdir(filepath.Join(appsdir, "test2"), os.FileMode(0777))

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

	assert.True(t, strings.HasSuffix(applicationsfound["test1"].BinaryPath, "bin"))
	assert.True(t, strings.HasSuffix(applicationsfound["test2"].BinaryPath, "test2"))
}

func TestListInstalledAppications_have_status_installed(t *testing.T) {

	appsdir := setupTestAppFolder()
	defer teardownTestAppFolder(appsdir)

	applicationsfound := ListInstalledAppications()

	assert.Equal(t, Installed, applicationsfound["test1"].Status)
	assert.Equal(t, Installed, applicationsfound["test2"].Status)
}

func TestAppicationsStatus(t *testing.T) {
	assert.Equal(t, "Available", Available.String())
	assert.Equal(t, "Installed", Installed.String())
	assert.Equal(t, "Activated", Activated.String())
	assert.Equal(t, "Remove", Remove.String())
	assert.Equal(t, "Status[100]", Status(100).String())
}
