package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitWithParameters(t *testing.T) {
	Init("codeshell_config_test.yaml")
	assert.Equal(t, viper.GetString("teststring"), "helloWorld", "failing load test config. Expecting teststring to be 'helloWorld'")
}

func TestInitNoParameters(t *testing.T) {
	Init()
}

func TestDefaultConfiguration(t *testing.T) {
	Init("codeshell_config_test.yaml")
	assert.Equal(t, viper.GetString("local.paths.applications"), "./apps/")
}
