package github

import (
	"github.com/praveenprem/sshauth/classes"
	"net/http"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"github.com/praveenprem/sshauth/config"
	"github.com/praveenprem/sshauth/logger"
	"github.com/praveenprem/sshauth/enums"
	"github.com/praveenprem/sshauth/github/classes"
)

/**
 * Package github
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 11/06/2018 21:39
 */

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
	var accessKeysList [] github.Key
	var result string
	if config.System_conf.Method == "team" {
		if config.System_conf.Admin_user == username {
			teamMembers := getTeamMembers(config.Github.Admin_role, config.Github)
			accessKeysList = keyCapture(teamMembers, publicKey, config)
		} else {
			teamMembers := getTeamMembers(config.Github.Default_role, config.Github)
			accessKeysList = keyCapture(teamMembers, publicKey, config)
		}
	}

	if config.System_conf.Method == "org" {
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

func getOrganizationMembers(role string, conf classes.GithubConf) []github.User {
	var members [] github.User
	var url = url_api + url_org + "/" + conf.Org + "/" + url_members + "?" + url_role + role + "&" + url_token + conf.Access_token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&members)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(5)
		}
	}
	return members
}

func getTeamMembers(role string, conf classes.GithubConf) []github.User {
	var members [] github.User
	teamId := getTeamId(conf.Org, conf.Access_token, conf.Team_name)
	var url = url_api + url_teams + "/" + strconv.Itoa(teamId) + "/" + url_members + "?" + url_role + role + "&" + url_token + conf.Access_token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&members)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(5)
		}
	}

	return members
}

func getTeamId(org string, token string, teamName string) int {
	var teams [] github.Team
	var url = url_api + url_org + "/" + org + "/" + url_teams + "?" + url_token + token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&teams)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(5)
		}

		for _, team := range teams {
			if team.Slug == teamName {
				return team.Id
			}
		}
	}

	logger.SimpleLogger(enums.ERROR, "unable to locate team \""+teamName+"\". Please check the configuration")
	os.Exit(61)

	return 0
}

func keyCapture(members [] github.User, publicKey string, conf classes.Conf) []github.Key {
	var keys []github.Key
	for _, member := range members {
		var userKeys [] github.Key
		var url = member.Url + "/" + url_keys + "?" + url_token + conf.Github.Access_token
		response, err := http.Get(url)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(70)
		} else {
			defer response.Body.Close()
			err := json.NewDecoder(response.Body).Decode(&userKeys)
			if err != nil {
				logger.SimpleLogger(enums.ERROR, err.Error())
				os.Exit(5)
			}
		}
		for _, k := range userKeys {
			if strings.Contains(k.Key, publicKey) {
				config.SendAlert(member.Login, "github.com")
			}

			keys = append(keys, k)
		}
	}

	return keys
}
