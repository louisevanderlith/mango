package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/pkg/enums"
)

type context struct {
	Users husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Users: husk.NewTable(new(User)),
	}

	createTestUsers()
}

func createTestUsers() {
	any := ctx.Users.Exists(husk.Everything())

	if any {
		return
	}

	user, err := NewUser("Admin", "admin@mango.avo")

	if err != nil {
		panic(err)
	}

	user.SecurePassword("Admin4v0")
	user.AddRole("Admin.APP", enums.Admin)

	ctx.Users.Create(user)
	defer ctx.Users.Save()
}
