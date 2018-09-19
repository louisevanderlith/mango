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

	"github.com/louisevanderlith/mango/util"
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

	var settings *Settings
	err := json.Unmarshal(content, settings)

	if err != nil {
		log.Print("loadSettings: ", err)
	}

	return settings
}

func (s *DomainSetting) SetupMux() error {
	lowType := strings.ToLower(s.Type)
	if s.Type == "subdomain" {
		return s.subdomainSetup()
	}

	if s.Type == "static" {
		return s.staticSetup()
	}

	msg := fmt.Sprintf("%s setting's Type '%s' was not found", s.Name, s.Type)
	return errors.New(msg)
}

func (s *DomainSetting) subdomainSetup() error {
	rawURL, err := util.GetServiceURL(s.Name, false)

	if err != nil {
		return err
	}

	if len(rawURL) == 0 {
		return errors.New("rawURL is empty")
	}

	vshost, err := url.Parse(rawURL)

	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(vshost)

	domainMux := http.NewServeMux()
	domainMux.HandleFunc("/", domainHandler(proxy))
	domains[s.Address] = domainMux

	log.Print(s.Name, " ", s.Address, " ", rawURL)

	return nil
}

func (s *DomainSetting) staticSetup() error {
	statMux := http.NewServeMux()
	fullPath := fmt.Sprintf("/static/%s/", s.Name)
	fullDir := http.Dir(fullPath)
	log.Printf("FullDIR: %s\n", fullDir)
	fs := http.FileServer(http.FileSystem(fullDir))

	statMux.Handle(fullPath, http.StripPrefix(fullPath, fs))
	statMux.Handle("/", fs)

	domains[s.Address] = statMux

	return nil
}
