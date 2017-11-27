package artifact

import (
	"github.com/astaxie/beego/orm"
	"github.com/louisevanderlith/mango/db"
)

type ArtifactContext struct {
	Upload *db.Set
}

var Ctx *ArtifactContext

func NewDatabase() {
	dbName := "Artifact.DB"

	registerModels()
	db.SyncDatabase(dbName)

	Ctx = &ArtifactContext{
		Upload: db.NewSet(Upload{}),
	}
}

func registerModels() {
	orm.RegisterModel(new(Upload))
}
