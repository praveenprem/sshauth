package logger

import (
	"os"
	"log"
	"log/syslog"
	"github.com/praveenprem/sshauth/enums"
)

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 14/06/2018 20:57
 */

const logFilePath = "/var/log/sshauth.log"

func SimpleLogger(level enums.Level, message string) {
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		if level == enums.INFO || level == enums.DEBUG {
			return
		}
		logFile, err := syslog.New(syslog.LOG_NOTICE, "sshauth")
		if err != nil {
			log.Fatalln(err)
			os.Exit(5)
		} else {
			defer logFile.Close()

			log.SetOutput(logFile)

			log.Printf("%s: %s\n", level, message)
		}
	} else {
		defer logFile.Close()

		log.SetOutput(logFile)

		log.Printf("%s: %s\n", level, message)
	}
}
