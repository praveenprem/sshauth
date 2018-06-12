package config

import (
	"os"
	"fmt"
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
		fmt.Println(file, err)
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
		//log.Fatalln("error: ", err)
		os.Exit(2)
	}

	return config
}

func Load() Conf {
	configuration = confParse()

	if configuration.System_conf.Service == "" ||
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
		if configuration.Mysql.Database == "" ||
			configuration.Mysql.Host == "" ||
			configuration.Mysql.Table == "" ||
			configuration.Mysql.User == "" ||
			configuration.Mysql.Password == ""  {
			exit("MySQL")
		}
	}

	return configuration
}

func exit(source string) {
	//log.Println("error: invalid configuration for source \""+ source +"\", please check the \"config.yml\" file")
	os.Exit(61)
}
