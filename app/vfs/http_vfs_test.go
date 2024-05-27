package vfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://ebamberg.github.io/codeshell/repository/applications.yaml
func TestChangeDirectory(t *testing.T) {
	fs, err := CreateHttpVFS("https://ebamberg.github.io/codeshell/")
	assert.Nil(t, err)
	cp, err := fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "https://ebamberg.github.io/codeshell/", cp)
	fs.Chdir("repository/")
	cp, err = fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "https://ebamberg.github.io/codeshell/repository/", cp)
	fs.Chdir("/absolutepath/")
	cp, err = fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "https://ebamberg.github.io/absolutepath/", cp)

}
