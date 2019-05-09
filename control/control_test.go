package control

import (
	"testing"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func TestApplication_GetPagedData_Page3(t *testing.T) {
	expect := 3
	actual, _ := getPageData("C5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}

func TestApplication_GetPagedData_Size5(t *testing.T) {
	expect := 5
	_, actual := getPageData("A5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}

func TestApplication_GetPagedData_Page26(t *testing.T) {
	expect := 26
	actual, _ := getPageData("Z5")

	if actual != expect {
		t.Errorf("Expecting %v, got %v", expect, actual)
	}
}

//Test Permissions...
func Test_tinyCtx_notoken_notAllowed_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "Test.API",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/abc", emptyMap)

	tiny, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/abc", "", roletype.Admin)

	if err != nil {
		t.Error(err)
	}

	allowed, err := tiny.allowed()

	if err != nil {
		t.Error(err)
	}

	if allowed {
		t.Errorf("Tiny was allowed. \n %v", tiny)
	}
}

func Test_tinyCtx_notoken_Allowed_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "Test.API",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Unknown

	ctrlmap.Add("/abc", emptyMap)

	tiny, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/abc", "", roletype.Unknown)

	if err != nil {
		t.Error(err)
	}

	allowed, err := tiny.allowed()

	if err != nil {
		t.Error(err)
	}

	if !allowed {
		t.Errorf("Tiny was not allowed. \n %v", tiny)
	}
}
