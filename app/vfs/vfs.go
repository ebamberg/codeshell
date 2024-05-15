package vfs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type VFSEntry struct {
	IsDir      bool
	Name       string
	Path       string
	filesystem *VFS
}

type VFS interface {
	Identifier() string
	List(path string) []VFSEntry
	Walk(path string, maxDepth int, callback func(entry VFSEntry)) error
	Chdir(path string) error
	Getwd() (string, error)
	Create(path string) (io.WriteCloser, error)
	Read(path string) (io.ReadCloser, error)
}

type ClosableReader struct {
	io.Reader
	io.Closer
}

type ClosableWriter struct {
	io.Reader
	io.Closer
}

var DefaultFilesystem VFS

func init() {
	DefaultFilesystem = LocalVFS{}
}

type LocalVFS struct {
}

func (this LocalVFS) Identifier() string {
	return "local-filesystem"
}

func (this LocalVFS) Create(path string) (io.WriteCloser, error) {
	out, err := os.Create(path)
	return out, err
}
func (this LocalVFS) Read(path string) (io.ReadCloser, error) {
	in, err := os.Open(path)
	return in, err
}

func (this LocalVFS) Chdir(path string) error {
	return os.Chdir(path)
}
func (this LocalVFS) Getwd() (string, error) {
	return os.Getwd()
}

func (this LocalVFS) List(path string) []VFSEntry {
	var result = make([]VFSEntry, 0)
	this.Walk(path, 0, func(entry VFSEntry) {
		result = append(result, entry)
	})
	return result
}

func (this LocalVFS) Walk(path string, maxDepth int, callback func(entry VFSEntry)) error {
	return filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.Count(path, string(os.PathSeparator)) > maxDepth {
			return fs.SkipDir
		}

		entry := VFSEntry{info.IsDir(), info.Name(), path, nil}
		callback(entry)
		return nil
	})
}
