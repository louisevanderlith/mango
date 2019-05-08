package control

import (
	"errors"
	"fmt"
	"log"
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
		for actPath, actMap := range m.mapping {
			if strings.Contains(path, actPath) {
				actionMap = actMap
				break
			}
		}
	}

	if actionMap == nil {
		panic(fmt.Sprintf("missing mapping for %s on %s", action, path))
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

	if strings.HasPrefix(path, "/static") || strings.HasPrefix(path, "/favicon") {
		return
	}

	url, token := removeToken(path)

	if token != "" {
		//localhost.org should resolve for development
		//TODO: write better function
		secure := false
		domain := ".localhost.org"

		if !strings.Contains(ctx.Request.Host, "localhost") {
			secure = true
			domain = "." + ctx.Request.Host
		}

		ctx.SetCookie("avosession", token, 0, "/", domain, secure, true)
	}

	if token == "" {
		authHeader := ctx.GetCookie("avosession")
		log.Println(authHeader)

		if len(authHeader) > 0 {
			token = strings.Split(authHeader, " ")[0]
		}
	}

	tiny, err := NewTinyCtx(m, ctx.Request.Method, url, token)
	securityURL, err := mango.GetServiceURL(m.GetInstanceID(), "Auth.APP", true)

	if err != nil {
		ctx.RenderMethodResult(RenderUnauthorized(err))
		return
	}

	req := ctx.Request
	moveURL := fmt.Sprintf("%s://%s%s", ctx.Input.Scheme(), req.Host, req.RequestURI)
	loginURL := buildLoginURL(securityURL, moveURL)

	allowed, err := tiny.allowed()

	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
		return
	}

	if err != nil || !allowed {
		// Redirect to login if not allowed.
		ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
		return
	}

	return
}

// FilterAPI is used to secure API services.
// When a user is not allowed to access a resource, they will get the Unauthorized Status.
func (m *ControllerMap) FilterAPI(ctx *context.Context) {
	path := ctx.Input.URL()

	if strings.HasPrefix(path, "/favicon") {
		return
	}

	url, token := removeToken(path)

	if token != "" {
		//localhost.org should resolve for development
		//TODO: write better function
		secure := false
		domain := ".localhost.org"

		if !strings.Contains(ctx.Request.Host, "localhost") {
			secure = true
			domain = "." + ctx.Request.Host
		}

		ctx.SetCookie("avosession", token, 0, "/", domain, secure, true)
	}

	if token == "" {
		authHeader := ctx.GetCookie("avosession")
		log.Println(authHeader)

		if len(authHeader) > 0 {
			token = strings.Split(authHeader, " ")[0]
		}
	}

	tiny, err := NewTinyCtx(m, ctx.Request.Method, url, token)

	if err != nil {
		ctx.RenderMethodResult(RenderUnauthorized(err))
		return
	}

	allowed, err := tiny.allowed()

	if err != nil {
		ctx.RenderMethodResult(RenderUnauthorized(err))
		return
	}

	if !allowed {
		ctx.RenderMethodResult(RenderUnauthorized(errors.New("no access")))
		return
	}

	return
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
