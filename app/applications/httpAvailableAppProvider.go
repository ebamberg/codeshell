package applications

import (
	"codeshell/vfs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type HttpAvailableApplicationProvider struct {
	repo vfs.VFS
}

// https://ebamberg.github.io/codeshell/repository/applications.yaml
func (this *HttpAvailableApplicationProvider) GetMapIndex() map[string][]Application {
	apps := make(map[string][]Application, 0)
	//	file, err := ioutil.ReadFile("C:\\dev\\src\\codeshell\\docs\\repository\\applications.yaml")
	file, err := this.repo.Read("applications.yaml")
	if err == nil {
		defer file.Close()
		buf, err := ioutil.ReadAll(file)

		if err == nil {
			err = yaml.Unmarshal(buf, &apps)
		}
		if err != nil {
			panic(err)
		}
	} else {
		log.Fatal(err)
	}
	return apps
}

func (this *HttpAvailableApplicationProvider) List() []Application {
	return FlattenMap(this.GetMapIndex())
}
