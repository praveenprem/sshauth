package gitlab

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 16/06/2018 15:20
 */

type User struct {
	Id int
	Name string
	Username string
	State string
	Avatar_url string
	Web_url string
	Created_at string
	Bio string
	Location string
	Skype string
	Linkedin string
	Twitter string
	Website_url string
	Organization string
	Last_sign_in_at string
	Confirmed_at string
	Last_activity_on string
	Email string
	Theme_id int
	Color_scheme_id int
	Projects_limit int
	Current_sign_in_at string
	Identities []Identity
	Can_create_group bool
	Can_create_project bool
	Two_factor_enabled bool
	External bool
	Is_admin bool
	Shared_runners_minutes_limit int
}

type Identity struct {
	Provider string
	Extern_uid string
}