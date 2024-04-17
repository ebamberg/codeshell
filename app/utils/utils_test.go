package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryExists(t *testing.T) {
	tmpDir := os.TempDir()
	assert.True(t, DirectoryExists(tmpDir))
	assert.False(t, DirectoryExists("/foobar"))
}

func TestDirectoryExists_when_pointing_to_file(t *testing.T) {
	tmpfile, _ := os.CreateTemp("", "unittest")
	result := DirectoryExists(tmpfile.Name())
	tmpfile.Close()
	assert.False(t, result)
}

func TestDirectoryExistsWithEmptyPath(t *testing.T) {
	assert.False(t, DirectoryExists(""))
}

// ENV VAR RELATED

func TestSetEnvVariableWithPathVariable(t *testing.T) {
	var testPath = "test/bin/"
	SetEnvVariable("path", testPath)
	var newPath = os.Getenv("PATH")
	assert.True(t, strings.HasPrefix(newPath, testPath) && len(newPath) > len(testPath))
	os.Setenv("CODESHELL_ORIGINAL_PATH", "")
}

func TestAppendPathVariable(t *testing.T) {
	var testPath = "test/bin/"
	AppendEnvPath(testPath)
	var newPath = os.Getenv("PATH")
	assert.True(t, strings.HasPrefix(newPath, testPath) && len(newPath) > len(testPath))
	os.Setenv("CODESHELL_ORIGINAL_PATH", "")
}

func TestResetEnvPath(t *testing.T) {
	originalPath := os.Getenv("PATH")
	var testPath = "test/bin/"
	SetEnvVariable("path", testPath)
	ResetEnvPath()
	var newPath = os.Getenv("PATH")
	assert.Equal(t, originalPath, newPath)
}

func TestResetEnvPathBeforeOriginalPathIsSet(t *testing.T) {
	originalPath := os.Getenv("PATH")
	os.Setenv("CODESHELL_ORIGINAL_PATH", "")
	ResetEnvPath()
	var newPath = os.Getenv("PATH")
	assert.Equal(t, originalPath, newPath)
}
