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
	assert.Equal(t, 3, len(profiles)) // default profile + 2 profiles that we created
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
	profile, exists := GetProfile("test1")
	assert.True(t, exists)
	assert.NotNil(t, profile)
	assert.Equal(t, "test profile", profile.Displayname)

}

func TestActivateProfile(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	activated := ActivateProfile("test1")
	assert.True(t, activated)
	assert.Equal(t, "hello world!", os.Getenv("TEST_CS_STRING"))

}

func TestActivateProfile_not_found(t *testing.T) {

	config.Init("codeshell_profiles_test.yaml")
	activated := ActivateProfile("gacklooma")

	assert.False(t, activated, fmt.Errorf("profile [gacklooma] not found"))

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
	os.MkdirAll(filepath.Join(appsdir, "java", "0.0.1", "bin"), os.FileMode(0777))
	os.MkdirAll(filepath.Join(appsdir, "maven", "0.0.1", "bin"), os.FileMode(0777))

	config.Set(CONFIG_KEY_APP_PATH, appsdir)
	return appsdir
}

func teardownTestAppFolder(appsdir string) {
	os.RemoveAll(appsdir)
}

func Test_ActivateApps(t *testing.T) {
	t.Skip("Problem: getInstalledApps relies on yaml config now..")
	config.Init("codeshell_profiles_test.yaml")
	testAppFolder := setupTestAppFolder()
	defer teardownTestAppFolder(testAppFolder)

	utils.ResetEnvPath()
	ActivateApps([]string{"java:0.0.1", "maven:0.0.1"}, true)
	path := config.GetString("Path")
	assert.True(t, strings.Contains(path, filepath.Join(testAppFolder, "java", "0.0.1", "bin")))
	assert.True(t, strings.Contains(path, filepath.Join(testAppFolder, "maven", "0.0.1", "bin")))

}
