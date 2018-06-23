package gitlab

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 16/06/2018 15:12
 */

type Members struct {
	Id int
	Name string
	Username string
	State string
	Avatar_url string
	Web_url string
	Access_level int
	Expires_at string
}