package gitlab

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 16/06/2018 15:08
 */

type Group struct {
	Id int
	Web_url string
	Name string
	Path string
	Description string
	Visibility string
	Lfs_enabled bool
	Avatar_url string
	Request_access_enabled bool
	Full_name string
	Full_path string
	Parent_id string
	Ldap_cn string
	Ldap_access string
}
