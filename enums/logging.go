package enums

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
