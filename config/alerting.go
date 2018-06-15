package config

import (
	"os"
	"bytes"
	"github.com/praveenprem/sshauth/classes"
	"encoding/json"
	"net/http"
	"github.com/praveenprem/sshauth/enums"
	"github.com/praveenprem/sshauth/logger"
)

const resourceLocal = "/tmp/sshAuthAlert.txt"

func SendAlert(user string, vender string) {
	conf := Load()
	if conf.Alerts == (classes.AlertConf{}) {
		logger.SimpleLogger(enums.WARNING, "Alerting skipped. No alerting configuration found")
	} else {
			if !checkLastKey(user) {
			if conf.Alerts.Slack != "" {
				slack(user, conf.System_conf.Name, conf.Alerts.Slack, vender)
			}
			if conf.Alerts.Hipchat != (classes.Hipchat{}) {
				hipChat(user, conf.System_conf.Name, conf.Alerts.Hipchat)
			}
			logNewAlert(user)
		} else {
			clearLast()
		}

	}
}

func isFileExist() bool {
	_, err := os.Stat(resourceLocal)
	if err == nil {
		return true
	} else {
		logger.SimpleLogger(enums.WARNING, err.Error())
		return false
	}
}

func checkLastKey(username string) bool {
	if !isFileExist() {
		return false
	} else {
		buffer := new(bytes.Buffer)

		file, err := os.Open(resourceLocal)
		if err != nil {
			logger.SimpleLogger(enums.WARNING, err.Error())
			return false
		}

		defer file.Close()

		buffer.ReadFrom(file)

		if buffer.String() == username {
			return true
		}
	}
	return false
}

func logNewAlert(username string) bool {
	logFile, err := os.OpenFile(resourceLocal, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.SimpleLogger(enums.WARNING, err.Error())
		return false
	}
	defer logFile.Close()

	logFile.Truncate(0)
	logFile.Seek(0, 0)
	logFile.Write([]byte(username))
	logFile.Sync()
	return true
}

func clearLast() {
	logFile, err := os.OpenFile(resourceLocal, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.SimpleLogger(enums.WARNING, err.Error())
	}
	defer logFile.Close()

	logFile.Truncate(0)
	logFile.Seek(0, 0)
	logFile.Sync()
}

func slack(user string, host string, url string, vender string) {
	var payload = classes.SlackPayloadBasic{}
	payload.Text = ">*New connection*\n```User: "+user+"\nHost:"+host+"\nService: "+vender+"```"

	body, err := json.Marshal(payload)
	if err != nil {
		logger.SimpleLogger(enums.WARNING, err.Error())
	} else {
		_, err = http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			logger.SimpleLogger(enums.WARNING, err.Error())
		}
	}

}

func hipChat(user string, host string, conf classes.Hipchat) {
//	TODO HipChat to be implemented
}
