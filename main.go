package main

import (
	"os"
	"github.com/praveenprem/sshauth/config"
	"github.com/praveenprem/sshauth/github"
	"fmt"
	"github.com/praveenprem/sshauth/enums"
	"github.com/praveenprem/sshauth/logger"
)

func main() {

	var keyChain string

	configs := config.Load()

	var user, pubKey string

	if len(os.Args) < 2 || len(os.Args) < 3 {
		logger.SimpleLogger(enums.ERROR,"arguments missing or not provided")
		os.Exit(22)
	} else if len(os.Args) > 4 {
		logger.SimpleLogger(enums.ERROR, "too many arguments provided")
		os.Exit(7)
	} else {
		user = os.Args[1]
		if user != configs.System_conf.Admin_user && user != configs.System_conf.Default_user {
			logger.SimpleLogger(enums.ERROR, "invalid user "+user)
			os.Exit(22)
		}

		if len(os.Args) == 4 {
			pubKey = os.Args[2] + " " + os.Args[3]
		} else {
			pubKey = os.Args[2]
		}
	}

	if configs.System_conf.Service == "github" {
		keyChain = github.Init(user, pubKey, configs)
	}

	fmt.Printf(keyChain)
}
