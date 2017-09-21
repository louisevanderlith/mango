package logic

import (
	"testing"

	uuid "github.com/nu7hatch/gouuid"
)

func dummyService(name string) Service {
	return Service{
		Environment: "dev",
		Name:        name,
		URL:         "http://127.0.01/" + name,
		Type:        "service"}
}

func TestAddService_ShouldCreateUUID(t *testing.T) {
	service := dummyService("Test.Service")

	result := AddService(service)

	if result == "" {
		t.Error("Didn't generate an UUID.")
	}
}

func TestGetService_AllowedCaller_ForDatabase_IsService(t *testing.T) {
	expect := "service"
	result := getService("Communication.DB", "dev", expect)

	if result.AllowedCaller != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCaller)
	}
}

func TestGetService_AllowedCaller_ForService_IsProxy(t *testing.T) {
	service := dummyService("Test.Service")

	AddService(service)

	expect := "proxy"
	result := getService("Test.Service", "dev", expect)

	if result.AllowedCaller != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCaller)
	}
}

func TestGetService_AllowedCaller_ForProxy_IsApplication(t *testing.T) {
	proxy := dummyService("Test.Proxy")
	proxy.Type = "proxy"

	AddService(proxy)

	expect := "application"
	result := getService("Test.Proxy", "dev", expect)

	if result.AllowedCaller != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCaller)
	}
}

func TestGetService_AllowedCaller_ForApplication_IsAll(t *testing.T) {
	app := dummyService("Test.App")
	app.Type = "application"

	AddService(app)

	result := getService("Test.App", "dev", "BLAH BLAH")
	expect := "*"

	if result.AllowedCaller != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCaller)
	}
}

func TestGetServicePath_SameEnv_ShouldFindService(t *testing.T) {
	requestor := dummyService("Test.Main")
	requestor.Type = "proxy"
	requestorID := AddService(requestor)

	api := dummyService("Test.Api")
	AddService(api)

	_, err := GetServicePath("Test.Api", requestorID)

	if err != nil {
		t.Error(err)
	}
}

func TestGetServicePath_DiffEnv_ShouldHaveError(t *testing.T) {
	requestor := dummyService("Test.Main")
	requestor.Type = "proxy"
	requestorID := AddService(requestor)

	api := dummyService("Test.Api")
	api.Environment = "uat"
	AddService(api)

	_, err := GetServicePath("Test.Api", requestorID)

	if err == nil {
		t.Error("Expecting an error message: Test.Api wasn't found for the requesting application")
	}
}

func TestGetServicePath_FakeRequestorID_ShouldHaveError(t *testing.T) {
	requestorID, _ := uuid.NewV4()

	api := dummyService("Test.Api")
	AddService(api)

	_, err := GetServicePath("Test.Api", requestorID.String())

	if err == nil {
		t.Error("Expecting an error message: Couldn't find an application with the given appID")
	}
}

func TestGetServicePath_ShouldFindDB(t *testing.T) {
	requestor := dummyService("Test.Main")
	requestorID := AddService(requestor)

	_, err := GetServicePath("Communication.DB", requestorID)

	if err != nil {
		t.Error(err)
	}
}
