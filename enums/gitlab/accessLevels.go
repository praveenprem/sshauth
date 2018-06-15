package gitlab

type level int

const (
	GUEST = 10
	REPORTER = 20
	DEVELOPER = 30
	MASTER = 40
	OWNER = 50
)

//func (level level) String() string {
//	l := []string{"GUEST", "REPORTER", "DEVELOPER", "MASTER", "OWNER"}
//
//	if level < GUEST || level > OWNER {
//		return "GUEST"
//	}
//
//	return l[level]
//}