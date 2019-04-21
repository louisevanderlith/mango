package control

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego/context"

	"github.com/louisevanderlith/mango"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

//ControllerMap is used to assign Priveliges to Actions
type ControllerMap struct {
	service *mango.Service
	mapping map[string]secure.ActionMap
}

func CreateControlMap(service *mango.Service) *ControllerMap {
	result := &ControllerMap{}
	result.service = service
	result.mapping = make(map[string]secure.ActionMap)

	return result
}

//Add is used to specify the permissions required for a controller's actions.
func (m *ControllerMap) Add(path string, actionMap secure.ActionMap) {
	m.mapping[path] = actionMap
}

//GetRequiredRole will return the RoleType required to access the 'path' and 'action'
func (m *ControllerMap) GetRequiredRole(path, action string) roletype.Enum {
	actionMap, hasCtrl := m.mapping[path]

	if !hasCtrl {
		return roletype.Unknown
	}

	roleType, hasAction := actionMap[action]

	if !hasAction {
		return roletype.Unknown
	}

	return roleType
}

//GetInstanceID returns the ID initially registered with the Service.
func (m *ControllerMap) GetInstanceID() string {
	return m.service.ID
}

//GetServiceName returns the Name initially registered with the Service
func (m *ControllerMap) GetServiceName() string {
	return m.service.Name
}

// FilterUI is used to secure web pages.
// When a user is not allowed to access a Page, they are redirected to secure.login
func (m *ControllerMap) FilterUI(ctx *context.Context) {
	path := ctx.Input.URL()

	if strings.HasPrefix(path, "/static") || strings.Contains(path, "favicon") {
		return
	}

	tiny := NewTinyCtx(m, ctx)

	if tiny.allowed() {
		return
	}

	instanceID := m.GetInstanceID()
	securityURL, err := mango.GetServiceURL(instanceID, "Auth.APP", true)

	if err != nil {
		return
	}

	req := ctx.Request
	moveURL := fmt.Sprintf("%s://%s%s", ctx.Input.Scheme(), req.Host, req.RequestURI)
	loginURL := buildLoginURL(securityURL, moveURL)

	// Redirect to login if not allowed.
	ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
}

// FilterAPI is used to secure API services.
// When a user is not allowed to access a resource, they will get the Unauthorized Status.
func (m *ControllerMap) FilterAPI(ctx *context.Context) {
	tiny := NewTinyCtx(m, ctx)

	if !tiny.allowed() {
		ctx.Abort(http.StatusUnauthorized, "User not authorized to access this content.")
	}
}

func buildLoginURL(securityURL, returnURL string) string {
	cleanReturn := removeQueries(returnURL)
	escURL := url.QueryEscape(cleanReturn)
	return fmt.Sprintf("%slogin?return=%s", securityURL, escURL)
}

func removeQueries(url string) string {
	idxOfQuery := strings.Index(url, "?")

	if idxOfQuery != -1 {
		url = url[:idxOfQuery]
	}

	return url
}
