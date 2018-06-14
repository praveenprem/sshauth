package github

import (
	"github.com/praveenprem/sshauth/classes"
	"net/http"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"github.com/praveenprem/sshauth/config"
)

const (
	url_api     = "https://api.github.com/"
	url_org     = "orgs"
	url_team    = "team"
	url_teams   = "teams"
	url_member  = "member"
	url_members = "members"
	url_keys    = "keys"
	url_token   = "access_token="
	url_role    = "role="
)

func Init(username string, publicKey string, config classes.Conf) string {
	var accessKeysList [] classes.GithubKey
	var result string
	if config.System_conf.Method == "team" {
		if config.System_conf.Admin_user == username {
			teamMembers := getTeamMembers(config.Github.Admin_role, config.Github)
			accessKeysList = keyCapture(teamMembers, publicKey, config)
		} else {
			teamMembers := getTeamMembers(config.Github.Default_role, config.Github)
			accessKeysList = keyCapture(teamMembers, publicKey, config)
		}
	} else if config.System_conf.Method == "org" {
		if config.System_conf.Admin_user == username {
			orgMembers := getOrganizationMembers(config.Github.Admin_role, config.Github)
			accessKeysList = keyCapture(orgMembers, publicKey, config)
		} else {
			orgMembers := getOrganizationMembers(config.Github.Default_role, config.Github)
			accessKeysList = keyCapture(orgMembers, publicKey, config)
		}
	}

	for _, key := range accessKeysList {
		result += key.Key + "\n"
	}
	return result
}

func getOrganizationMembers(role string, conf classes.GithubConf) []classes.GithubUser {
	var members [] classes.GithubUser
	var url = url_api + url_org + "/" + conf.Org + "/" + url_members + "?" + url_role + role + "&" + url_token + conf.Access_token
	response, err := http.Get(url)
	if err != nil {
		//log.Printf("ERROR: %s\n", err)
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&members)
		if err != nil {
			//log.Printf("ERROR: %s\n", err)
			os.Exit(5)
		}
	}
	return members
}

func getTeamMembers(role string, conf classes.GithubConf) []classes.GithubUser {
	var members [] classes.GithubUser
	teamId := listTeams(conf.Org, conf.Access_token, conf.Team_name)
	var url = url_api + url_teams + "/" + strconv.Itoa(teamId) + "/" + url_members + "?" + url_role + role + "&" + url_token + conf.Access_token
	response, err := http.Get(url)
	if err != nil {
		//log.Printf("ERROR: %s\n", err)
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&members)
		if err != nil {
			//log.Printf("ERROR: %s\n", err)
			os.Exit(5)
		}
	}

	return members
}

func listTeams(org string, token string, teamName string) int {
	var teams [] classes.GithubTeam
	var url = url_api + url_org + "/" + org + "/" + url_teams + "?" + url_token + token
	response, err := http.Get(url)
	if err != nil {
		//log.Printf("ERROR: %s\n", err)
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&teams)
		if err != nil {
			//log.Printf("ERROR: %s\n", err)
			os.Exit(5)
		}

		for _, team := range teams {
			if team.Slug == teamName {
				return team.Id
			}
		}
	}
	
	//log.Printf("ERROR: %s\n", err)
	os.Exit(61)

	return 0
}

func keyCapture(members [] classes.GithubUser, publicKey string, conf classes.Conf) []classes.GithubKey {
	var keys []classes.GithubKey
	for _, member := range members {
		var userKeys [] classes.GithubKey
		var url = member.Url + "/" + url_keys + "?" + url_token + conf.Github.Access_token
		response, err := http.Get(url)
		if err != nil {
			//log.Printf("ERROR: %s\n", err)
			os.Exit(70)
		} else {
			defer response.Body.Close()
			err := json.NewDecoder(response.Body).Decode(&userKeys)
			if err != nil {
				//log.Printf("ERROR: %s\n", err)
				os.Exit(5)
			}
		}
		for _, k := range userKeys {
			if strings.Contains(k.Key, publicKey) {
				config.SendAlert(member.Html_url, publicKey)
			}

			keys = append(keys, k)
		}
	}

	return keys
}
