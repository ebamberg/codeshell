package applications

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type HttpAvailableApplicationProvider struct {
}

func (this *HttpAvailableApplicationProvider) GetMapIndex() map[string][]Application {
	apps := make(map[string][]Application, 0)
	file, err := ioutil.ReadFile("C:\\dev\\src\\codeshell\\docs\\repository\\applications.yaml")

	if err == nil {
		err = yaml.Unmarshal(file, &apps)
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
