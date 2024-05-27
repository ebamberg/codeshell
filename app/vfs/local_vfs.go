package vfs

import (
	"errors"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type LocalVFS struct {
	rootpath   string
	currentdir string
}

func CreateLocalVFS(base_url string) (*LocalVFS, error) {
	bu, err := url.Parse(base_url)
	if err == nil {
		return FromUrlLocalFile(bu)
	} else {
		return &LocalVFS{}, err
	}

}
func FromUrlLocalFile(base_url *url.URL) (*LocalVFS, error) {
	if base_url.Scheme != "file" {
		return &LocalVFS{}, errors.New("unsupported url-scheme. url-scheme has to be file.")
	}
	return &LocalVFS{rootpath: base_url.Path}, nil
}

func (this *LocalVFS) Identifier() string {
	return "local-filesystem"
}

func (this *LocalVFS) Create(path string) (io.WriteCloser, error) {
	out, err := os.Create(path)
	return out, err
}
func (this *LocalVFS) Read(path string) (io.ReadCloser, error) {
	in, err := os.Open(path)
	return in, err
}

func (this *LocalVFS) Chdir(path string) error {
	return os.Chdir(path)
}
func (this *LocalVFS) Getwd() (string, error) {
	return os.Getwd()
}

func (this *LocalVFS) List(path string, maxDepth int) []VFSEntry {
	var result = make([]VFSEntry, 0)
	this.Walk(path, maxDepth, func(entry VFSEntry) {
		result = append(result, entry)
	})
	return result
}

func (this *LocalVFS) Walk(path string, maxDepth int, callback func(entry VFSEntry)) error {
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

func (this *LocalVFS) resolve(path string) string {
	var result string
	if this.rootpath == "" {
		return path
	} else {
		result = filepath.Join(this.rootpath, this.currentdir, path)
	}
	return result
}
