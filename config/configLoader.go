package config

import (
	"os"
	"bytes"
	"gopkg.in/yaml.v2"
	"github.com/praveenprem/sshauth/classes"
	"github.com/praveenprem/sshauth/enums"
	"github.com/praveenprem/sshauth/logger"
)

type Conf = classes.Conf

var configuration Conf

func init() {
	configuration = Conf{}
}

func loadFile() string {
	configFile := "/etc/sshauth/config.yml"
	buffer := new(bytes.Buffer)

	file, err := os.Open(configFile)

	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(2)
	}

	defer file.Close()

	buffer.ReadFrom(file)

	return buffer.String()
}

func confParse() Conf {
	configFileContent := loadFile()
	var config Conf

	err := yaml.Unmarshal([]byte(configFileContent), &config)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(2)
	}

	return config
}

func Load() Conf {
	configuration = confParse()

	if configuration.System_conf.Service == "" ||
		configuration.System_conf.Name == "" ||
		configuration.System_conf.Admin_user == "" ||
		configuration.System_conf.Default_user == "" {
		exit("System")
	}

	if configuration.System_conf.Service == "github" {
		if configuration.System_conf.Method == "" ||
			configuration.Github.Access_token == "" ||
			configuration.Github.Admin_role == "" ||
			configuration.Github.Default_role == "" ||
			configuration.Github.Org == "" {
			exit("GitHub")
		} else {
			if configuration.System_conf.Method == "team" {
				if configuration.Github.Team_name == "" {
					exit("GitHub")
				}
			}
		}
	}

	if configuration.System_conf.Service == "gitlab" {
		if configuration.System_conf.Method == "" ||
			configuration.Gitlab.Access_token == "" ||
			configuration.Gitlab.Admin_role == "" ||
			configuration.Gitlab.Default_role == "" ||
			configuration.Gitlab.Org_url == "" {
			exit("GitLab")
		} else {
			if configuration.System_conf.Method == "group" {
				if configuration.Gitlab.Group_name == "" {
					exit("GitLab")
				} else {
					if configuration.Gitlab.Inherit_permission.Admin_user && configuration.Gitlab.Inherit_permission.Admin_stack == "" {
						exit("GitLab -> inherit_permission")
					} else {
						if configuration.Gitlab.Inherit_permission.Admin_user && (configuration.Gitlab.Inherit_permission.Admin_stack != classes.UP &&
							configuration.Gitlab.Inherit_permission.Admin_stack != classes.DOWN) {
							exit("GitLab -> admin_stack")
						}
					}
					if configuration.Gitlab.Inherit_permission.Default_user && configuration.Gitlab.Inherit_permission.Default_stack == "" {
						exit("GitLab -> inherit_permission")
					} else {
						if configuration.Gitlab.Inherit_permission.Default_user && (configuration.Gitlab.Inherit_permission.Default_stack != classes.UP &&
							configuration.Gitlab.Inherit_permission.Default_stack != classes.DOWN) {
							exit("GitLab -> default_stack")
						}
					}
				}
			}
		}
	}

	if configuration.System_conf.Service == "database" {
		if configuration.Sql.Database == "" ||
			configuration.Sql.Host == "" ||
			configuration.Sql.Table == "" ||
			configuration.Sql.User == "" ||
			configuration.Sql.Password == "" {
			exit("SQLConf")
		}
	}

	if configuration.Alerts.Hipchat != (classes.Hipchat{}) {
		if configuration.Alerts.Hipchat.Url == "" ||
			configuration.Alerts.Hipchat.Token == "" {
			exit("HipChat")
		}
	}

	return configuration
}

func exit(source string) {
	logger.SimpleLogger(enums.ERROR, "Invalid configuration for source \""+source+"\", please check the \"config.yml\" file")
	os.Exit(61)
}
