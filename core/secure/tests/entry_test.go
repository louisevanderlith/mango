package secure_test

import (
	"testing"

	"github.com/louisevanderlith/mango/util/control"
	"github.com/nu7hatch/gouuid"

	"github.com/louisevanderlith/mango/core/secure"
)

func getFakeApp() control.Application {
	instID, _ := uuid.NewV4()
	result := control.Application{
		InstanceID: instID.String(),
		IP:         "127.0.0.1",
		Location:   "-26.1496832, 28.035481599999997",
		Name:       "Entry.TEST",
	}

	return result
}

func TestRegistration_Good_Pass(t *testing.T) {
	r := secure.Registration{
		App:            getFakeApp(),
		Email:          "joe@fake.com",
		Name:           "Joe",
		Password:       "w34k###",
		PasswordRepeat: "w34k###",
	}

	rec, err := secure.Register(r)

	if err != nil {
		t.Error(err)
	}

	data := rec.Data().(*secure.User)

	if len(data.Roles) == 0 {
		t.Error("No Roles")
	}
}

func TestRegistration_Bad_Fail(t *testing.T) {
	r := secure.Registration{
		App:      getFakeApp(),
		Name:     "Joe",
		Password: "w34k###",
	}

	rec, err := secure.Register(r)

	if err == nil {
		t.Log("no error found error")
		t.Error(rec)
	}
}

func TestRegistration_Ugly_Fail(t *testing.T) {
	r := secure.Registration{
		App:            getFakeApp(),
		Email:          "joe%fake.com",
		Name:           "Joe",
		Password:       "w34k!c$651d",
		PasswordRepeat: "w34k!c$651d",
	}

	rec, err := secure.Register(r)

	if err == nil {
		t.Log("no error found error")
		t.Error(rec)
	}
}

func TestLogin_Good_Pass(t *testing.T) {
	r := secure.Registration{
		App:            getFakeApp(),
		Email:          "joe@fake.com",
		Name:           "Joe",
		Password:       "w34k###",
		PasswordRepeat: "w34k###",
	}

	_, err := secure.Register(r)

	if err != nil {
		t.Error(err)
		return
	}

	if err != nil && err.Error() != "email already in use" {
		t.Error(err)
		return
	}

	authreq := secure.Authentication{
		App:      getFakeApp(),
		Email:    "joe@fake.com",
		Password: "w34k###",
	}

	_, err = secure.Login(authreq)

	if err != nil {
		t.Error(err)
	}
}

func TestLogin_Bad_Fail(t *testing.T) {

}

func TestLogin_Ugly_Fail(t *testing.T) {

}
