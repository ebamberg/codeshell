package applications

import (
	"codeshell/config"

	"github.com/spf13/viper"
)

type LocalInstalledApplicationProvider struct {
}

func (this *LocalInstalledApplicationProvider) GetMapIndex() map[string][]Application {
	apps := make(map[string][]Application, 0)
	err := viper.UnmarshalKey(config.CONFIG_KEY_APPLICATIONS_INSTALLED, &apps)
	if err != nil {
		panic(err)
	}
	return apps
}

func (this *LocalInstalledApplicationProvider) List() []Application {
	return FlattenMap(this.GetMapIndex())
}
