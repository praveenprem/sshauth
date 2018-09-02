package gitlab

import (
	"github.com/praveenprem/sshauth/config"
	"github.com/praveenprem/sshauth/gitlab/classes"
	"net/http"
	"github.com/praveenprem/sshauth/logger"
	"github.com/praveenprem/sshauth/enums"
	"os"
	"encoding/json"
	"strings"
	"strconv"
	"github.com/praveenprem/sshauth/classes"
	"github.com/praveenprem/sshauth/alert"
)

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 15/06/2018 19:48
 */

const (
	URL_API     = "/api/v4/"
	URL_GROUPS  = "groups"
	URL_MEMBERS = "members"
	URL_KEYS    = "keys"
	URL_USERS   = "users"
	URL_TOKEN   = "private_token="
)

func Init(username string, publicKey string, conf config.Conf) string {
	var keyList []gitlab.Key
	var result string
	if conf.System_conf.Method == "group" {
		if username == conf.System_conf.Admin_user {
			keyList = getKeys(
				getGroupMembers(conf.Gitlab.Admin_role, conf.Gitlab.Inherit_permission.Admin_user,
					conf.Gitlab.Inherit_permission.Admin_stack, conf),
				nil, publicKey, conf)
		} else {
			keyList = getKeys(
				getGroupMembers(conf.Gitlab.Default_role, conf.Gitlab.Inherit_permission.Default_user,
					conf.Gitlab.Inherit_permission.Default_stack, conf),
				nil, publicKey, conf)
		}
	}
	if conf.System_conf.Method == "org" {
		if username == conf.System_conf.Admin_user {
			keyList = getKeys(
				nil, getOrganisationUsers(conf.Gitlab.Admin_role, conf), publicKey, conf)
		} else {
			keyList = getKeys(
				nil, getOrganisationUsers(conf.Gitlab.Default_role, conf), publicKey, conf)
		}
	}

	for _, key := range keyList {
		result += key.Key + "\n"
	}

	return result
}

func getOrganisationUsers(role string, conf config.Conf) []gitlab.User {
	var tmpUsers []gitlab.User
	var users []gitlab.User
	var url = conf.Gitlab.Org_url + URL_API + URL_USERS + "/?" + URL_TOKEN + conf.Gitlab.Access_token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(70)
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&tmpUsers)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(5)
		} else {
			if strings.ToLower(role) == "admin" {
				for _, user := range tmpUsers {
					if user.Is_admin {
						users = append(users, user)
					}
				}
			} else if strings.ToLower(role) == "members" {
				for _, user := range tmpUsers {
					if !user.Is_admin {
						users = append(users, user)
					}
				}
			} else if strings.ToLower(role) == "all" {
				users = tmpUsers
			} else {
				logger.SimpleLogger(enums.ERROR, "invalid role declared in the Gitlab configuration")
				os.Exit(61)
			}
		}
	}
	return users
}

func getGroupMembers(role string, permission bool, stack string, conf config.Conf) []gitlab.Members {
	var tmpUsers []gitlab.Members
	var users []gitlab.Members
	var groupId = getGroupId(conf)
	var LEVEL = int(0)

	var url = conf.Gitlab.Org_url + URL_API + URL_GROUPS + "/" + strconv.Itoa(groupId) + "/" + URL_MEMBERS + "/?" + URL_TOKEN + conf.Gitlab.Access_token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
		os.Exit(70)
	} else {
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&tmpUsers)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
			os.Exit(5)
		}

	}

	if len(tmpUsers) > 0 {
		switch strings.ToUpper(role) {
		case "OWNER":
			LEVEL = gitlab.OWNER
		case "MASTER":
			LEVEL = gitlab.MASTER
		case "DEVELOPER":
			LEVEL = gitlab.DEVELOPER
		case "REPORTER":
			LEVEL = gitlab.REPORTER
		case "GUEST":
			LEVEL = gitlab.GUEST
		default:
			logger.SimpleLogger(enums.ERROR, "invalid role declared in the Gitlab configuration")
			os.Exit(61)
		}

		if LEVEL > 0 {
			users = filterGroupMembers(
				tmpUsers,
				LEVEL,
				permission,
				stack)
		}
	}
	return users
}

func filterGroupMembers(members []gitlab.Members, level int, permission bool, stack string) []gitlab.Members {
	var m [] gitlab.Members
	if permission {
		if stack == classes.UP {
			for _, user := range members {
				if user.Access_level <= level {
					m = append(m, user)
				}
			}
		} else {
			for _, user := range members {
				if user.Access_level >= level {
					m = append(m, user)
				}
			}
		}
	} else {
		for _, user := range members {
			if user.Access_level == level {
				m = append(m, user)
			}
		}
	}
	return m
}

func getGroupId(conf config.Conf) int {
	var groups []gitlab.Group
	var url = conf.Gitlab.Org_url + URL_API + URL_GROUPS + "?" + URL_TOKEN + conf.Gitlab.Access_token
	response, err := http.Get(url)
	if err != nil {
		logger.SimpleLogger(enums.ERROR, err.Error())
	} else {
		defer response.Body.Close()

		err := json.NewDecoder(response.Body).Decode(&groups)
		if err != nil {
			logger.SimpleLogger(enums.ERROR, err.Error())
		}

		for _, group := range groups {
			if group.Name == conf.Gitlab.Group_name {
				return group.Id
			}
		}
	}

	logger.SimpleLogger(enums.ERROR, "unable to locate group \""+conf.Gitlab.Group_name+"\". Please check the configuration")
	os.Exit(61)

	return 0
}

func getKeys(groupMembers []gitlab.Members, orgUsers []gitlab.User, publicKey string, conf config.Conf) []gitlab.Key {
	var keys []gitlab.Key
	if groupMembers != nil {
		for _, member := range groupMembers {
			var userKeys []gitlab.Key
			var url = conf.Gitlab.Org_url + URL_API + URL_USERS + "/" + strconv.Itoa(member.Id) + "/" + URL_KEYS + "?" + URL_TOKEN + conf.Gitlab.Access_token
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
					slack.SendAlert(member.Username, conf.Gitlab.Org_url)
				}

				keys = append(keys, k)
			}
		}
	} else if orgUsers != nil {
		for _, member := range orgUsers {
			var userKeys []gitlab.Key
			var url = conf.Gitlab.Org_url + URL_API + URL_USERS + "/" + strconv.Itoa(member.Id) + "/" + URL_KEYS + "?" + URL_TOKEN + conf.Gitlab.Access_token
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
					slack.SendAlert(member.Username, conf.Gitlab.Org_url)
				}

				keys = append(keys, k)
			}
		}
	} else {
		logger.SimpleLogger(enums.ERROR, "unable to fetch keys, group members and org users are null")
		os.Exit(22)
	}
	return keys
}
