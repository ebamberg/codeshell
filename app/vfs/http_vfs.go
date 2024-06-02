package vfs

import (
	"codeshell/output"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type HttpVFS struct {
	baseurl     *url.URL
	currentPath *url.URL
}

func CreateHttpVFS(base_url string) (*HttpVFS, error) {
	bu, err := url.Parse(base_url)
	if err == nil {
		return FromHttpVFSUrl(bu)
	} else {
		return &HttpVFS{}, err
	}

}

func FromHttpVFSUrl(base_url *url.URL) (*HttpVFS, error) {
	if base_url.Scheme != "http" && base_url.Scheme != "https" {
		return &HttpVFS{}, errors.New("unsupported url-scheme. url-scheme has to be http or https.")
	}
	return &HttpVFS{baseurl: base_url, currentPath: base_url}, nil
}

func (this *HttpVFS) Identifier() string {
	return "http-based-filesystem"
}

func (this *HttpVFS) Create(path string) (io.WriteCloser, error) {
	return nil, errors.New("not yet supported")
}
func (this *HttpVFS) Read(path string) (io.ReadCloser, error) {
	filepath, err := this.currentPath.Parse(path)
	if err == nil {
		resp, err := http.Get(filepath.String())
		if err == nil {
			return resp.Body, nil
		}
	}
	return nil, err
}

func (this *HttpVFS) Chdir(path string) error {
	newPath, err := this.currentPath.Parse(path)
	if err == nil {
		this.currentPath = newPath
		return nil
	} else {
		return errors.New("not yet supported")
	}
}
func (this *HttpVFS) Getwd() (string, error) {
	return this.currentPath.String(), nil
}

func (this *HttpVFS) Exists(path string) bool {
	output.Errorln("fucntion Exists not yet supported for HTTP based filesystems")
	return false
}

func (this *HttpVFS) List(path string, maxDepth int) []VFSEntry {
	return make([]VFSEntry, 0)
}

func (this *HttpVFS) Walk(path string, maxDepth int, callback func(entry VFSEntry)) error {
	return errors.New("not yet supported")
}
