package enums

/**
 * Package enums
 * Project sshauth
 * Created by Praveen Premaratne
 * Created on 14/06/2018 22:32
 */

type Level int

const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
)

func (level Level) String() string {
	l := []string{
		"ERROR",
		"WARNING",
		"INFO",
		"DEBUG"}
	if level < ERROR || level > DEBUG{
		return "WARNING"
	}

	return l[level]
}
