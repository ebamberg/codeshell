package applications

var available = map[string][]Application{
	"eclipse-jee": []Application{
		{
			Id:          "eclipse-jee",
			DisplayName: "Eclipse for Java and WebDev",
			Status:      Available,
			Version:     "2024-03",
			Source: appInstallationSource{
				Size:             545459269,
				Url:              "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-jee-2024-03-R-win32-x86_64.zip&r=1",
				IgnoreRootFolder: true,
			}},
	},
	"eclipse-cpp": []Application{
		{
			Id:          "eclipse-cpp",
			DisplayName: "Eclipse for C/C++ developer",
			Status:      Available,
			Version:     "2024-03",
			Source: appInstallationSource{
				Size:             0,
				Url:              "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-cpp-2024-03-R-win32-x86_64.zip&r=1",
				IgnoreRootFolder: true,
			}},
	},
	"maven": []Application{
		{
			Id:          "maven",
			DisplayName: "Apache Maven",
			Status:      Available,
			Version:     "3.9.6",
			Source: appInstallationSource{
				Size:             9513253,
				Url:              "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.zip",
				IgnoreRootFolder: true,
				EnvVars: map[string]string{
					"MAVEN_HOME": "{{.Path}}",
				},
			}},
	},
	"npp": []Application{
		{
			Id:          "npp",
			DisplayName: "Notepad++",
			Status:      Available,
			Version:     "8.6.7",
			Source: appInstallationSource{
				Size: 5998909,
				Url:  "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.7/npp.8.6.7.portable.x64.zip",
			}},
	},
	"python": []Application{
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.12",
			Source: appInstallationSource{
				Size: 0,
				Url:  "https://www.python.org/ftp/python/3.12.3/python-3.12.3-embed-amd64.zip",
			}},
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.11",
			Source: appInstallationSource{
				Size: 0,
				Url:  "https://www.python.org/ftp/python/3.11.9/python-3.11.9-embed-amd64.zip",
			}},
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.9",
			Source: appInstallationSource{
				Size: 0,
				Url:  "https://www.python.org/ftp/python/3.9.13/python-3.9.13-embed-amd64.zip",
			}},
	},
}
