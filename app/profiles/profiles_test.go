package profiles

import (
	"codeshell/config"
	"fmt"
	"os"
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
	assert.Equal(t, "def", profiles["test1"].Applications)

	assert.Equal(t, "ghi", profiles["test2"].Applications)

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
