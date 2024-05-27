package vfs

import (
	"errors"
	"io"
	"net/url"
)

type VFSEntry struct {
	IsDir      bool
	Name       string
	Path       string
	filesystem *VFS
}

type VFS interface {
	Identifier() string
	List(path string, maxDepth int) []VFSEntry
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
	DefaultFilesystem = &LocalVFS{}
}

func FromUrlString(urlString string) (VFS, error) {
	base_url, err := url.Parse(urlString)
	if err == nil {
		return FromUrl(base_url)
	} else {
		return &LocalVFS{}, err
	}
}

func FromUrl(base_url *url.URL) (VFS, error) {
	if base_url.Scheme == "http" || base_url.Scheme == "https" {
		return FromHttpVFSUrl(base_url)
	} else if base_url.Scheme == "file" {
		return FromUrlLocalFile(base_url)
	}

	return &LocalVFS{}, errors.New("unsupported url-scheme. scheme is " + base_url.Scheme)
}
