package profiles

import (
	"codeshell/config"
	"codeshell/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProfiles(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	profiles, err := getAllProfiles()
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, profiles)
	assert.Equal(t, 2, len(profiles))
	assert.Equal(t, "test profile", profiles["test1"].Displayname)
	assert.Equal(t, []string([]string{"java", "maven"}), profiles["test1"].Applications)

	assert.Equal(t, []string([]string{"ghi"}), profiles["test2"].Applications)

}

func Test_GetAllProfiles_not_provided_empty_string(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	profiles, err := getAllProfiles()
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, profiles)
	assert.Equal(t, "", profiles["test2"].Displayname)
}

func Test_GetAllProfiles_envVars_list(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	profiles, err := getAllProfiles()
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, profiles)

	assert.Equal(t, 3, len(profiles["test1"].EnvVars))

	fmt.Print(profiles["test1"].EnvVars)

	keys := []string{"test_cs_string", "test_cs_2", "test_cs_3"}
	expected := []string{"hello world!", "foobar", "123"}

	for i, key := range keys {
		assert.Equal(t, expected[i], profiles["test1"].EnvVars[key])

	}

	assert.Equal(t, 0, len(profiles["test2"].EnvVars))
}

func TestGetProfile(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	profile, err := GetProfile("test1")
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, "test profile", profile.Displayname)

}

func TestActivateProfile(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	err := ActivateProfile("test1")
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, "hello world!", os.Getenv("TEST_CS_STRING"))

}

func TestActivateProfile_not_found(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	err := ActivateProfile("gacklooma")

	assert.NotErrorIs(t, err, fmt.Errorf("profile [gacklooma] not found"))

}

func Test_GetAllProfiles_application_list(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	profiles, err := getAllProfiles()
	if err != nil {
		t.Log(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, profiles)

	assert.Equal(t, 2, len(profiles["test1"].Applications))
}

const CONFIG_KEY_APP_PATH = "local.paths.applications"

func setupTestAppFolder() string {
	appsdir, _ := os.MkdirTemp("", "apps")
	os.Mkdir(filepath.Join(appsdir, "java"), os.FileMode(0777))
	os.Mkdir(filepath.Join(appsdir, "java", "bin"), os.FileMode(0777))
	os.Mkdir(filepath.Join(appsdir, "maven"), os.FileMode(0777))
	os.Mkdir(filepath.Join(appsdir, "maven", "bin"), os.FileMode(0777))

	config.Set(CONFIG_KEY_APP_PATH, appsdir)
	return appsdir
}

func teardownTestAppFolder(appsdir string) {
	os.RemoveAll(appsdir)
}

func Test_ActivateApps(t *testing.T) {
	config.Init("codeshell_profiles_test.yaml")
	testAppFolder := setupTestAppFolder()
	defer teardownTestAppFolder(testAppFolder)

	utils.ResetEnvPath()
	ActivateApps([]string{"java", "maven"})
	path := config.GetString("Path")
	assert.True(t, strings.Contains(path, filepath.Join(testAppFolder, "java", "bin")))
	assert.True(t, strings.Contains(path, filepath.Join(testAppFolder, "maven", "bin")))

}
