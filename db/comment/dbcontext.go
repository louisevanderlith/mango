package comment

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type CommentContext struct {
	Comment *db.Set
}

var Ctx *CommentContext

func NewDatabase() {
	dbName := "Comment.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &CommentContext{
		Comment: db.NewSet(Comment{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Comment))
}
