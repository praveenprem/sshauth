package config

import (
	"os"
	"bytes"
	"gopkg.in/yaml.v2"
	"github.com/praveenprem/sshauth/classes"
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
		//log.Fatalf("ERROR: %s\n", err)
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
		//log.Fatalf("ERROR: %s\n", err)
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

	if configuration.System_conf.Service == "database" {
		if configuration.Sql.Database == "" ||
			configuration.Sql.Host == "" ||
			configuration.Sql.Table == "" ||
			configuration.Sql.User == "" ||
			configuration.Sql.Password == ""  {
			exit("SQLConf")
		}
	}

	return configuration
}

func exit(source string) {
	//log.Fatalln("ERROR: Invalid configuration for source \""+ source +"\", please check the \"config.yml\" file")
	os.Exit(61)
}
