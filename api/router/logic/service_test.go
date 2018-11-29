package logic

import (
	"testing"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	uuid "github.com/nu7hatch/gouuid"
)

func dummyService(name string) *util.Service {
	return util.NewService(enums.DEV.String(), name, enums.API)
	/*  util.Service{
	Environment: enums.LIVE,
	Name:        name,
	URL:         "http://127.0.01/" + name,
	Type:        enums.API}*/
}

func TestAddService_ShouldCreateUUID(t *testing.T) {
	service := dummyService("Test.Service")

	result, err := AddService(service)

	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("didn't generate UUID")
	}
}

func TestGetService_AllowedCaller_ForApplication_IsAll(t *testing.T) {
	app := dummyService("Test.App")
	app.Type = enums.APP

	AddService(app)

	result := getService("Test.App", enums.DEV, app.Type)
	expect := enums.ANY

	if result.AllowedCaller != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCaller.String())
	}
}

func TestGetServicePath_SameEnv_ShouldFindService(t *testing.T) {
	requestor := dummyService("Test.Main")
	requestor.Type = enums.APP
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	api := dummyService("Test.Api")
	AddService(api)

	_, err = GetServicePath("Test.Api", requestorID, false)

	if err != nil {
		t.Error(err)
	}
}

func TestGetServicePath_DiffEnv_ShouldHaveError(t *testing.T) {
	requestor := dummyService("Test.Main")
	requestor.Type = enums.APP
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	api := dummyService("Test.Api")
	api.Environment = enums.UAT
	AddService(api)

	_, err = GetServicePath("Test.Api", requestorID, false)

	if err == nil {
		t.Error("Expecting an error message: Test.Api wasn't found for the requesting application")
	}
}

func TestGetServicePath_FakeRequestorID_ShouldHaveError(t *testing.T) {
	requestorID, _ := uuid.NewV4()

	api := dummyService("Test.Api")
	AddService(api)

	_, err := GetServicePath("Test.Api", requestorID.String(), false)

	if err == nil {
		t.Error("Expecting an error message: Couldn't find an application with the given appID")
	}
}

func TestGetServicePath_SameService__CantCallSelf_ShouldHaveError(t *testing.T) {
	requestor := dummyService("Test.API")
	requestor.Type = enums.API
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	_, err = GetServicePath("Test.Api", requestorID, false)

	if err == nil {
		t.Error("Expecting 'Test.Api wasn't found for the requesting application'")
	}
}
