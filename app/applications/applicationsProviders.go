package applications

import (
	"codeshell/config"
	"codeshell/vfs"
)

var localAppProvider ApplicationProvider

var providers []ApplicationProvider

//{
// &HttpAvailableApplicationProvider{},
//	&InternalAvailableApplicationProvider{},
//	&LocalInstalledApplicationProvider{},
//	localAppProvider,
//}

func GetProviders() []ApplicationProvider {
	if providers == nil {
		initProviders()
	}
	return providers
}

func initProviders() {
	providers = make([]ApplicationProvider, 0, 2)
	repo_url := config.GetString(config.CONFIG_KEY_REPO_APP_URL)
	var availableProvider ApplicationProvider
	if repo_url != "" {
		repo_fs, err := vfs.FromUrlString(repo_url)
		if err == nil {
			availableProvider = &HttpAvailableApplicationProvider{repo: repo_fs}

		} else {
			panic(err)
		}
	} else {
		availableProvider = &InternalAvailableApplicationProvider{}
	}
	providers = append(providers, availableProvider)
	localAppProvider = &LocalInstalledApplicationProvider{}
	providers = append(providers, localAppProvider)

}

// ************************ Often used Predicated ********************

func IsInstalled(a Application) bool {
	return a.Status == Installed
}

// ************************ Helper Functions **************************

func FlattenMap(appMap map[string][]Application) []Application {
	v := make([]Application, 0, len(appMap))

	for _, values := range appMap {
		for _, value := range values {
			v = append(v, value)
		}
	}
	return v
}
