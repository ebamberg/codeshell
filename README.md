# codeshell
Codeshell - useful environment for developers

# Getting started

Download the codeshell release for you operating system from the latest release in this repository:

for example
https://github.com/ebamberg/codeshell/releases/download/v0.1.5/codeshell_Windows_x86_64.zip

unpack the archive in any folder. The archive contains a codeshell.exe.

```shell
codeshell.exe help
```
will print a list of all commands.

When no command is passed then codeshell start in interactive shell mode same as you would start codeshell with
```shell
codeshell shell
```

## configuration file
When starting the exe without a configuration-file codeshell starts on a default values.

you can define the location of the configuration file with

```shell 
codeshell shell --config <path_name_of_config_yaml_file>
```

# Install/uninstall applications

list all installed applications as well as available application. 
```shell
apps list --all
```

An application is available if it is not locally uninstalled but codeshell is able to install it.
To install an application we have to pass the id as shown in the application list.

example

```shell
apps install maven
```

If multiple version of the same application is available we can specify the version we want to install

```shell
apps install maven:3.9
```

or 
```shell
apps install python:3.11
```

to uninstall an application we can use

```shell
apps uninstall maven
```

## activating a profile

After installation an application is not __activated__.

To activate an application we have to add it to a profile in the config.yaml file.

We can define multiple different profiles in the configuration file each containing
a different set of application.
By activating a profile we are activating all applications defined in the profile.

An activated application is add the PATH of your OS (for the current shell) and necessary
configurations like setting environment variables is done. The activation is just valid as long
as the profile is activated. We can switch over to another profile at any time activating a different
set of applications

```shell
profile activate default
```

To list configured profiles we can use
```shell
profiles list 
```

### example profile configuration
```yaml
profiles:
    default:
        applications:
            - npp:8.6.7
            - maven:3.9.6
            - eclipse-jee:2024-03
        autoinstallapps: "true"
        displayname: default
        envvars:
            foobar: foobar
            hello: "world"
            test_cs_string: hello world!
        id: default
    ml:
        applications:
            - python:3.12
        autoinstallapps: "true"
        displayname: MachineLearning
        id: ml
```

# the Configuration file

```yaml
applications:
    installed:                              # list of installed applications
        eclipse-jee:
            - id: eclipse-jee
              displayname: Eclipse for Java and WebDev
              path: c:\dev\codeshell\apps\eclipse-jee\2024-03
              binarypath: c:\dev\codeshell\apps\eclipse-jee\2024-03
              status: 2
              version: 2024-03
              envvars: {}
        maven:
            - id: maven
              displayname: Apache Maven
              path: c:\dev\codeshell\apps\maven\3.9.6
              binarypath: c:\dev\codeshell\apps\maven\3.9.6
              status: 2
              version: 3.9.6
              envvars: {}
        npp:
            - id: npp
              displayname: Notepad++
              path: c:\dev\codeshell\apps\npp\8.6.7
              binarypath: c:\dev\codeshell\apps\npp\8.6.7
              status: 2
              version: 8.6.7
              envvars: {}
        python:
            - id: python
              displayname: Python
              path: c:\dev\codeshell\apps\python\3.12
              binarypath: c:\dev\codeshell\apps\python\3.12
              status: 2
              version: "3.12"
              envvars: {}
local:
    paths:
        applications: c:\dev\codeshell\apps     # path where the application are downloaded and installed
profiles:
    default:                                    # profile with id="default"
        applications:                           # list of application assigned to this profiles. when activating this profiles. this apps are activated
            - npp:8.6.7
            - maven:3.9.6
            - eclipse-jee:2024-03
        autoinstallapps: "true"                 # when "true" the codeshell automatically installs apps of the profile if the are not already installed
        displayname: default
        envvars:                                # profile related env variable that are set. each application has it's own envvars section for application specific envvars     
            test_cs_2: foobar
            test_cs_3: "123"
            test_cs_string: hello world!
        id: default
    ml:                                         # profile with id="ml"
        applications:
            - python:3.12
        autoinstallapps: "true"
        displayname: MachineLearning
        id: ml
repositories:
    applications:
        url: https://ebamberg.github.io/codeshell/repository/   # https server to look up available applications. under the address we expect an applications.yaml file with information where to download and how to run an application
terminal:
    style:
        title: Codeshell                        # the title bar when starting codeshell.

```

