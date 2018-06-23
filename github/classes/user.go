package github

/**
 * Package github
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 16/06/2018 14:54
 */

type User struct {
	Login string
	Id int
	Node_id string
	Avatar_url string
	Gravatar_id string
	Url string
	Html_url string
	Followers_url string
	Following_url string
	Gists_url string
	Starred_url string
	Subscriptions_url string
	Organizations_url string
	Repos_url string
	Events_url string
	Received_events_url string
	Type string
	Site_admin bool
}
