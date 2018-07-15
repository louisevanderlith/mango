package secure

type context struct {
	LoginTraces loginTracesTable
	Roles       rolesTable
	Users       usersTable
}

var ctx context

func NewContext() {
	ctx = context{
		LoginTraces: NewLoginTracesTable(),
		Roles:       NewRolesTable(),
		Users:       NewUsersTable(),
	}
}
