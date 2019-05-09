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
func (m *ControllerMap) GetRequiredRole(path, action string) (roletype.Enum, error) {
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
		return roletype.Unknown, fmt.Errorf("missing mapping for %s on %s", action, path)
	}

	roleType, hasAction := actionMap[strings.ToUpper(action)]

	if !hasAction {
		return roletype.Unknown, nil
	}

	return roleType, nil
}

//GetInstanceID returns the ID initially registered with the Service.
func (m *ControllerMap) GetInstanceID() string {
	return m.service.ID
}

//GetServiceName returns the Name initially registered with the Service
func (m *ControllerMap) GetServiceName() string {
	return m.service.Name
}

//GetPublicKeyPath return the Location of the public key file registered with the Service.
func (m *ControllerMap) GetPublicKeyPath() string {
	return m.service.PublicKey
}

// FilterUI is used to secure web pages.
// When a user is not allowed to access a Page, they are redirected to secure.login
func (m *ControllerMap) FilterUI(ctx *context.Context) {
	path := ctx.Request.URL.RequestURI()
	action := ctx.Request.Method

	log.Printf("FilterUI: %s | %s\n", path, action)

	if strings.HasPrefix(path, "/static") || strings.HasPrefix(path, "/favicon") {
		return
	}

	requiredRole, err := m.GetRequiredRole(path, action)

	if err != nil {
		log.Println(err)
		//Missing Mapping, the user doesn't have access to the application, and must request it.
		sendToSubscription(ctx, m.GetInstanceID())
		return
	}

	if requiredRole == roletype.Unknown {
		return
	}

	url, token := removeToken(path)

	if token == "" {
		authHeader := ctx.Request.Header["Authorization"]

		if len(authHeader) > 0 {
			token = strings.Split(authHeader[0], " ")[0]
		} else {
			log.Println("no authorization found")
			sendToLogin(ctx, m.GetInstanceID())
			return
		}
	}

	tiny, err := NewTinyCtx(m.GetServiceName(), ctx.Request.Method, url, token, requiredRole, m.GetPublicKeyPath())

	if err != nil {
		log.Println(err)
		sendToLogin(ctx, m.GetInstanceID())
		return
	}

	allowed, err := tiny.allowed()

	if err != nil || !allowed {
		log.Println(err)
		sendToLogin(ctx, m.GetInstanceID())
	}
}

// FilterAPI is used to secure API services.
// When a user is not allowed to access a resource, they will get the Unauthorized Status.
func (m *ControllerMap) FilterAPI(ctx *context.Context) {
	path := ctx.Request.URL.RequestURI()
	action := ctx.Request.Method

	if strings.HasPrefix(path, "/favicon") {
		return
	}

	requiredRole, err := m.GetRequiredRole(path, action)

	if err != nil {
		//Missing Mapping, the user doesn't have access to the application
		ctx.RenderMethodResult(RenderUnauthorized(err))
		return
	}

	if requiredRole == roletype.Unknown {
		return
	}

	authHeader := ctx.Request.Header["Authorization"]

	if len(authHeader) == 0 {
		err := errors.New("no authorization header provided")
		ctx.RenderMethodResult(RenderUnauthorized(err))
		return
	}

	token := strings.Split(authHeader[0], " ")[0]
	tiny, err := NewTinyCtx(m.GetServiceName(), action, path, token, requiredRole, m.GetPublicKeyPath())

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
	}
}

func sendToLogin(ctx *context.Context, instanceID string) {
	securityURL, err := mango.GetServiceURL(instanceID, "Auth.APP", true)

	if err != nil {
		ctx.RenderMethodResult(err)
		return
	}

	req := ctx.Request
	moveURL := fmt.Sprintf("%s://%s%s", ctx.Input.Scheme(), req.Host, req.RequestURI)
	loginURL := buildLoginURL(securityURL, moveURL)

	ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
}

func sendToSubscription(ctx *context.Context, instanceID string) {
	securityURL, err := mango.GetServiceURL(instanceID, "Auth.APP", true)

	if err != nil {
		ctx.RenderMethodResult(err)
		return
	}

	subcribeURL := buildSubscribeURL(securityURL)

	ctx.Redirect(http.StatusTemporaryRedirect, subcribeURL)
}

func buildLoginURL(securityURL, returnURL string) string {
	cleanReturn := removeQueries(returnURL)
	escURL := url.QueryEscape(cleanReturn)
	return fmt.Sprintf("%slogin?return=%s", securityURL, escURL)
}

func buildSubscribeURL(securityURL string) string {
	return fmt.Sprintf("%ssubscribe", securityURL)
}

func removeQueries(url string) string {
	idxOfQuery := strings.Index(url, "?")

	if idxOfQuery != -1 {
		url = url[:idxOfQuery]
	}

	return url
}
