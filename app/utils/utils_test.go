package utils

import (
	"os"
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
