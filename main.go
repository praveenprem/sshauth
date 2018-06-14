package main

import (
	"os"
	"github.com/praveenprem/sshauth/config"
	"github.com/praveenprem/sshauth/github"
	"fmt"
	"strings"
)

//TODO Add system logging support

func main() {

	var keyChain string

	configs := config.Load()

	var user, pubKey string

	if len(os.Args) < 2 || len(os.Args) < 3 {
		//log.Printf("ERROR: arguments missing or not provided")
		os.Exit(22)
	} else if len(os.Args) > 4 {
		//log.Printf("ERROR: too many arguments provided")
		os.Exit(7)
	} else {
		user = os.Args[1]
		if user != configs.System_conf.Admin_user && user != configs.System_conf.Default_user {
			//log.Printf("ERROR: invalid user %s", user)
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

	fmt.Printf(strings.TrimSpace(keyChain))
}
