package classes

/**
 * Package classes
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 13/06/2018 22:49
 */

type AlertConf struct {
	Hipchat Hipchat
	Slack   string
}

type SlackPayloadBasic struct {
	Text string `json:"text"`
}

type SlackPayloadPetty struct {
	Text string `json:"text"`
	Username string `json:"username"`
	Mrkdwn bool `json:"mrkdwn"`

}

type Hipchat struct {
	Url string
	Token string
}