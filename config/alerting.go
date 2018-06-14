package config

import (
	"os"
	"bytes"
	"github.com/praveenprem/sshauth/classes"
	"encoding/json"
	"net/http"
)

const resourceLocal = "/tmp/sshAuthAlert.txt"

func SendAlert(user string, publicKey string) {
	conf := Load()
	if conf.Alerts == (classes.AlertConf{}) {
		//log.Println("WARN: Alerting skipped. No alerting configuration found")
	} else {
		if !checkLastKey(publicKey) {
			//	TODO Add HTTP trigger for alert with payload
			//	TODO Revise the logic here
			slack(user, conf.System_conf.Name, conf.Alerts.Slack)
			logNewAlert(publicKey)
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
		//log.Printf("WARN: %s\n", err)
		return false
	}
}

func checkLastKey(publicKey string) bool {
	if !isFileExist() {
		//log.Printf("WARN: %s\n", err)
		return false
	} else {
		buffer := new(bytes.Buffer)

		file, err := os.Open(resourceLocal)
		if err != nil {
			//log.Printf("WARN: %s\n", err)
			return false
		}

		defer file.Close()

		buffer.ReadFrom(file)

		if buffer.String() == publicKey {
			return true
		}
	}
	return false
}

func logNewAlert(publicKey string) bool {
	logFile, err := os.OpenFile(resourceLocal, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		//log.Printf("WARN: %s\n", err)
		return false
	}
	defer logFile.Close()

	logFile.Truncate(0)
	logFile.Seek(0, 0)
	logFile.Write([]byte(publicKey))
	logFile.Sync()
	return true
}

func slack(user string, host string, url string) {
	var payload = classes.AlertPayload{}
	payload.Text = "User: "+user+" has SSH in to ```"+host+"```"

	body, err := json.Marshal(payload)
	if err != nil {
		//log.Printf("WARN: %s\n", err)
	} else {
		_, err = http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			//log.Printf("WARN: %s\n", err)
		}
	}

}

func clearLast() {
	logFile, err := os.OpenFile(resourceLocal, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		//log.Printf("WARN: %s\n", err)
	}
	defer logFile.Close()

	logFile.Truncate(0)
	logFile.Seek(0, 0)
	logFile.Sync()
}