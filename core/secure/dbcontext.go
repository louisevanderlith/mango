package secure

type context struct {
	LoginTraces loginTracesTable
	Roles       rolesTable
	Users       usersTable
}

var ctx context

func init() {
	ctx = context{
		LoginTraces: NewLoginTracesTable(),
		Roles:       NewRolesTable(),
		Users:       NewUsersTable(),
	}
}
