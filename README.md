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



