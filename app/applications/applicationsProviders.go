package applications

var localAppProvider = &LocalInstalledApplicationProvider{}

var Providers = []ApplicationProvider{
	&InternalAvailableApplicationProvider{},
	//	&LocalInstalledApplicationProvider{},
	localAppProvider,
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
