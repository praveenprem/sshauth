package classes

type Github struct {
	Access_token string
	Admin_role string
	Api_url string
	Default_role string
	Key_url string
	Org string
	Team_name string
}

type GithubTeam struct {
	Name string
	Id int
	Node_id string
	Slug string
	Description string
	Privacy string
	Url string
	Members_url string
	Repositories_url string
	Permission string
}

type GithubUser struct {
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

type GithubKey struct {
	Id int
	Key string
}