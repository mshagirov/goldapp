package config

const (
	UserFilter  = "(objectClass=PosixAccount)"
	GroupFilter = "(objectClass=PosixGroup)"
)

var (
	UsrCols = []string{"Username", "uid", "Name", "Group"}
	UsrAttr = map[string]string{
		"uid":       "Username",
		"uidNumber": "uid",
		"cn":        "Name",
		"gidNumber": "Group",
	}
	UsrColsWidth = []int{15, 5, 15, 10, 25}

	GrpCols = []string{"Name", "gid", "dn"}
	GrpAttr = map[string]string{
		"cn":        "Name",
		"gidNumber": "gid",
		"dn":        "dn",
	}
	GrpColsWidth = []int{15, 5, 25}
)
