package secure

type context struct {
	Users usersTable
}

var ctx context

func init() {
	ctx = context{
		Users: NewUsersTable(),
	}
}
