package gitlab

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
	Parent_id int
	Ldap_cn string
	Ldap_access string
}
