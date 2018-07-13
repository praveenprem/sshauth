package gitlab

/**
 * Package gitlab
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 16/06/2018 15:06
 */

type level int

const (
	GUEST = 10
	REPORTER = 20
	DEVELOPER = 30
	MASTER = 40
	OWNER = 50
)