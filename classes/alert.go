package classes

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