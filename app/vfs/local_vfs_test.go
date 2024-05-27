package vfs

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalFSResolveNoBase(t *testing.T) {
	fs := LocalVFS{}
	assert.Equal(t, "relative/", fs.resolve("relative/"))
}

func TestLocalFSResolveRelative(t *testing.T) {
	fs := LocalVFS{rootpath: "/codeshell/root/"}
	assert.Equal(t, filepath.Join("/codeshell/root/", "relative/"), fs.resolve("relative/"))
}

func TestLocalFSResolveRelativeWithCurrent(t *testing.T) {
	fs := LocalVFS{rootpath: "/codeshell/root/", currentdir: "work/"}
	assert.Equal(t, filepath.Join("/codeshell/root/", "work", "relative/"), fs.resolve("relative/"))
}

func TestChangeDirectoryLocalFS(t *testing.T) {
	t.Skip("not yet implemented")
	fs, err := CreateLocalVFS("file://codeshell/root/")
	assert.Nil(t, err)
	cp, err := fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "/", cp)
	fs.Chdir("repository/")
	cp, err = fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "/repository/", cp)
	fs.Chdir("/absolutepath/")
	cp, err = fs.Getwd()
	assert.Nil(t, err)
	assert.Equal(t, "/absolutepath/", cp)

}

func TestChangeDirectoryLocalFS_internal(t *testing.T) {
	fs, err := CreateLocalVFS("file://codeshell/root/")
	assert.Nil(t, err)
	assert.NotNil(t, fs)
}
