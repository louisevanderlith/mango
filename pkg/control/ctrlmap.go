package control

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego/context"

	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/enums"
)

//Action map[:action]:roleType
type ActionMap map[string]enums.RoleType

//ControllerMap is used to assign Priveliges to Actions
type ControllerMap struct {
	service *util.Service
	mapping map[string]ActionMap
}

func CreateControlMap(service *util.Service) *ControllerMap {
	result := &ControllerMap{}
	result.service = service
	result.mapping = make(map[string]ActionMap)

	return result
}

// AddControllerMap is used to specify the permissions required for a controller's actions.
func (m *ControllerMap) Add(path string, actionMap ActionMap) {
	m.mapping[path] = actionMap
}

func (m *ControllerMap) GetRequiredRole(path, action string) enums.RoleType {
	actionMap, hasCtrl := m.mapping[path]

	if !hasCtrl {
		return enums.Unknown
	}

	roleType, hasAction := actionMap[action]

	if !hasAction {
		return enums.Unknown
	}

	return roleType
}

func (m *ControllerMap) GetInstanceID() string {
	return m.service.ID
}

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
	log.Printf("Tiny: %v allowed %#v\n", tiny.allowed(), tiny)

	if tiny.allowed() {
		return
	}

	instanceID := m.GetInstanceID()
	securityURL, err := util.GetServiceURL(instanceID, "Secure.API", true)

	if err != nil {
		log.Printf("FilterUI Failed: %+v\n", err)
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
	return fmt.Sprintf("%sv1/login?return=%s", securityURL, escURL)
}

func removeQueries(url string) string {
	idxOfQuery := strings.Index(url, "?")

	if idxOfQuery != -1 {
		url = url[:idxOfQuery]
	}

	return url
}
