package classes

type GitLabConf struct {
	Access_token string
	Admin_role string
	Default_role string
	Group_name string
	Org_url string
	Inherit_permission Permission
}

type Permission struct {
	Admin_user    bool
	Admin_stack   string
	Default_user  bool
	Default_stack string

}

const (
	UP = "up"
	DOWN = "down"
)