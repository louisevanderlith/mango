package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/nu7hatch/gouuid"
)

type LoginTrace struct {
	Location        string `hsk:"null;size(128)"`
	IP              string `hsk:"null;size(50)"`
	Allowed         bool   `hsk:"default(true)"`
	InstanceID      uuid.UUID
	ApplicationName string `hsk:"size(20)"`
	TraceType
}

func (o LoginTrace) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func getRegistrationTrace(r Registration) LoginTrace {
	return LoginTrace{
		Allowed:         true,
		ApplicationName: r.App.Name,
		InstanceID:      r.App.InstanceID,
		IP:              r.App.IP,
		Location:        r.App.Location,
		TraceType:       TraceRegister,
	}
}

func getLoginTrace(r Authentication, passed bool) LoginTrace {
	trace := TraceLogin

	if !passed {
		trace = TraceFail
	}

	return LoginTrace{
		Allowed:         passed,
		ApplicationName: r.App.Name,
		InstanceID:      r.App.InstanceID,
		IP:              r.App.IP,
		Location:        r.App.Location,
		TraceType:       trace,
	}
}
