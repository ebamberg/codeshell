package applications

var available = map[string][]Application{
	"eclipse-jee": []Application{
		{
			Id:          "eclipse-jee",
			DisplayName: "Eclipse for Java and WebDev",
			Status:      Available,
			Version:     "2024-03",
			source: appInstallationSource{
				size:             545459269,
				url:              "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-jee-2024-03-R-win32-x86_64.zip&r=1",
				ignoreRootFolder: true,
			}},
	},
	"eclipse-cpp": []Application{
		{
			Id:          "eclipse-cpp",
			DisplayName: "Eclipse for C/C++ developer",
			Status:      Available,
			Version:     "2024-03",
			source: appInstallationSource{
				size:             0,
				url:              "https://www.eclipse.org/downloads/download.php?file=/technology/epp/downloads/release/2024-03/R/eclipse-cpp-2024-03-R-win32-x86_64.zip&r=1",
				ignoreRootFolder: true,
			}},
	},
	"maven": []Application{
		{
			Id:          "maven",
			DisplayName: "Apache Maven",
			Status:      Available,
			Version:     "3.9.6",
			source: appInstallationSource{
				size:             9513253,
				url:              "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.zip",
				ignoreRootFolder: true,
				envVars: map[string]string{
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
			source: appInstallationSource{
				size: 5998909,
				url:  "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.7/npp.8.6.7.portable.x64.zip",
			}},
	},
	"python": []Application{
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.12",
			source: appInstallationSource{
				size: 0,
				url:  "https://www.python.org/ftp/python/3.12.3/python-3.12.3-embed-amd64.zip",
			}},
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.11",
			source: appInstallationSource{
				size: 0,
				url:  "https://www.python.org/ftp/python/3.11.9/python-3.11.9-embed-amd64.zip",
			}},
		{
			Id:          "python",
			DisplayName: "Python",
			Status:      Available,
			Version:     "3.9",
			source: appInstallationSource{
				size: 0,
				url:  "https://www.python.org/ftp/python/3.9.13/python-3.9.13-embed-amd64.zip",
			}},
	},
}
