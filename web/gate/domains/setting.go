package domains

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/louisevanderlith/mango/pkg"
)

type DomainSetting struct {
	Address string
	Name    string
	Type    string
}

type Settings []*DomainSetting

func loadSettings() *Settings {
	dbConfPath := util.FindFilePath("domains.json", "conf")
	content := util.GetFileContent(dbConfPath)

	settings := &Settings{}
	err := json.Unmarshal(content, settings)

	if err != nil {
		log.Print("loadSettings: ", err)
	}

	return settings
}

func (s *DomainSetting) SetupMux(instanceID string) (http.Handler, error) {
	lowType := strings.ToLower(s.Type)

	if lowType == "subdomain" {
		return s.subdomainSetup(instanceID)
	}

	if lowType == "static" {
		return s.staticSetup()
	}

	msg := fmt.Sprintf("%s setting's Type '%s' was not found", s.Name, s.Type)
	return nil, errors.New(msg)
}

func (s *DomainSetting) subdomainSetup(instanceID string) (http.Handler, error) {
	rawURL, err := util.GetServiceURL(instanceID, s.Name, false)

	if err != nil {
		return nil, err
	}

	vshost, err := url.Parse(rawURL)

	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(vshost)

	domainMux := http.NewServeMux()
	domainMux.HandleFunc("/", proxy.ServeHTTP)

	log.Printf("Name:%s\tAddr:%s\n", s.Name, s.Address)

	return domainMux, nil
}

func (s *DomainSetting) staticSetup() (http.Handler, error) {
	statMux := http.NewServeMux()
	fullPath := fmt.Sprintf("/static/%s/", s.Name)
	fullDir := http.Dir(fullPath)

	log.Printf("FullDIR: %s\n", fullDir)

	fs := http.FileServer(http.FileSystem(fullDir))

	statMux.Handle(fullPath, http.StripPrefix(fullPath, fs))
	statMux.Handle("/", fs)

	return statMux, nil
}
