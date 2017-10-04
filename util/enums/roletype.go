package enums

type RoleType int

const (
	Admin RoleType = iota
	Owner
	User
)

var roleTypes = [...]string{
	"Admin",
	"Owner",
	"User"}

func (r RoleType) String() string {
	return roleTypes[r]
}
