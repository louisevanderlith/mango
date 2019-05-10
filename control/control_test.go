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

func TestRemoveToken_MustRemoveParsableToken(t *testing.T) {
	expected := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJUCI6IiIsIkxvY2F0aW9uIjoiIiwiVXNlcktleSI6IjE1NTY3MDYzOTJgMCIsIlVzZXJSb2xlcyI6eyJBZG1pbi5BUFAiOjAsIlJvdXRlci5BUEkiOjB9LCJVc2VybmFtZSI6IkFkbWluIn0.w4VlzJXwRnmK7P0lUeqw1xR1eLYr9Fdne6e8LBwdim0mCHVSVlKCsFonMQQ097rirpFBXNyGI8QNpLIBfafwLh2zz9YwviUyMkEX400P_8UO2lCiYGtLVe-tf9Ndz9xsnZUnFOz8nuOuRlWGQbcGw8zrE20rmEGxU-daOb9RVnGqtrN3ANzwj9StgI67IQtkCR27a40n9mdi7R9Sx7U4Bc2cz5oAT44-JNzESXNmlrbWWLxLnRs4TTGLrMS5fhzO_v_SgyuZaAj6kOTjEiGiHCbIqQAX4GlzmYw1OEGu8r-Sr3nxGcm91svt7ddmEdc1dVlMMj-3xy3oHBWWAFULtw"
	redirect := "https://admin.localhost/?access_token=" + expected
	url, token := removeToken(redirect)

	if token == "" {
		t.Error("token not found")
	}

	if token != expected {
		t.Errorf("Expected %s, got %s", expected, token)
	}

	if url != "https://admin.localhost/" {
		t.Error("url incorrect", url)
	}
}

//Test Permissions...
func Test_tinyCtx_ValidToken_Allowed_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "Admin.APP",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
		PublicKey:     "../../secure/db/sign_rsa.pub",
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/", emptyMap)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJUCI6IiIsIkxvY2F0aW9uIjoiIiwiVXNlcktleSI6IjE1NTY3MDYzOTJgMCIsIlVzZXJSb2xlcyI6eyJBZG1pbi5BUFAiOjAsIlJvdXRlci5BUEkiOjB9LCJVc2VybmFtZSI6IkFkbWluIn0.w4VlzJXwRnmK7P0lUeqw1xR1eLYr9Fdne6e8LBwdim0mCHVSVlKCsFonMQQ097rirpFBXNyGI8QNpLIBfafwLh2zz9YwviUyMkEX400P_8UO2lCiYGtLVe-tf9Ndz9xsnZUnFOz8nuOuRlWGQbcGw8zrE20rmEGxU-daOb9RVnGqtrN3ANzwj9StgI67IQtkCR27a40n9mdi7R9Sx7U4Bc2cz5oAT44-JNzESXNmlrbWWLxLnRs4TTGLrMS5fhzO_v_SgyuZaAj6kOTjEiGiHCbIqQAX4GlzmYw1OEGu8r-Sr3nxGcm91svt7ddmEdc1dVlMMj-3xy3oHBWWAFULtw"
	tiny, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/", token, roletype.Admin, ctrlmap.GetPublicKeyPath())

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

func Test_tinyCtx_ValidToken_NotAllowedMapping_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "Theme.API",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
		PublicKey:     "../../secure/db/sign_rsa.pub",
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/", emptyMap)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJUCI6IiIsIkxvY2F0aW9uIjoiIiwiVXNlcktleSI6IjE1NTY3MDYzOTJgMCIsIlVzZXJSb2xlcyI6eyJBZG1pbi5BUFAiOjAsIlJvdXRlci5BUEkiOjB9LCJVc2VybmFtZSI6IkFkbWluIn0.w4VlzJXwRnmK7P0lUeqw1xR1eLYr9Fdne6e8LBwdim0mCHVSVlKCsFonMQQ097rirpFBXNyGI8QNpLIBfafwLh2zz9YwviUyMkEX400P_8UO2lCiYGtLVe-tf9Ndz9xsnZUnFOz8nuOuRlWGQbcGw8zrE20rmEGxU-daOb9RVnGqtrN3ANzwj9StgI67IQtkCR27a40n9mdi7R9Sx7U4Bc2cz5oAT44-JNzESXNmlrbWWLxLnRs4TTGLrMS5fhzO_v_SgyuZaAj6kOTjEiGiHCbIqQAX4GlzmYw1OEGu8r-Sr3nxGcm91svt7ddmEdc1dVlMMj-3xy3oHBWWAFULtw"
	tiny, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/", token, roletype.Admin, ctrlmap.GetPublicKeyPath())

	if err != nil {
		t.Error(err)
	}

	allowed, err := tiny.allowed()

	if err == nil {
		t.Error("expected 'mapping'")
	}

	if allowed {
		t.Errorf("Tiny was allowed. \n %v", tiny)
	}
}

func Test_tinyCtx_NoToken_NotAllowed_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "Admin.APP",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
		PublicKey:     "../../secure/db/sign_rsa.pub",
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Unknown

	ctrlmap.Add("/", emptyMap)

	_, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/", "", roletype.Unknown, ctrlmap.GetPublicKeyPath())

	if err == nil {
		t.Error("expecting 'invalid token'")
	}
}

func Test_tinyCtx_NoToken_Allowed_HasApplicationRole_Admin(t *testing.T) {
	srvc := &mango.Service{
		ID:            "X",
		Name:          "WWW.APP",
		AllowedCaller: enums.APP,
		Type:          enums.API,
		Environment:   enums.DEV,
		PublicKey:     "../../secure/db/sign_rsa.pub",
	}

	ctrlmap := CreateControlMap(srvc)
	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.Unknown

	ctrlmap.Add("/", emptyMap)

	tiny, err := NewTinyCtx(ctrlmap.GetServiceName(), "GET", "/", "", roletype.Unknown, ctrlmap.GetPublicKeyPath())

	if err != nil {
		t.Error(err)
		return
	}

	allowed, err := tiny.allowed()

	if err != nil {
		t.Error(err)
	}

	if !allowed {
		t.Errorf("Tiny was not allowed. \n %v", tiny)
	}
}
