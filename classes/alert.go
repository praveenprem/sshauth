package classes

type AlertConf struct {
	Hipchat Hipchat
	Slack   string
}

type AlertPayload struct {
	Text string `json:"text"`
}

type Hipchat struct {
	Url string
	Token string
}