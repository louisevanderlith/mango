package secure

import (
	"github.com/louisevanderlith/mango/util/control"
	"github.com/nu7hatch/gouuid"
)

type AuthRequest struct {
	ApplicationName string
	InstanceID      uuid.UUID
	Email           string
	Password        []byte
	IP              string
	Location        string
}

func (req AuthRequest) GetApplication() *control.Application {
	return control.NewApplication(req.ApplicationName, req.InstanceID)
}
