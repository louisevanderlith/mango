package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/louisevanderlith/mango/util/enums"
)

type Service struct {
	ID            string
	Name          string
	URL           string
	Version       int
	Environment   enums.Environment
	AllowedCaller enums.ServiceType
	Type          enums.ServiceType
}

func NewService(env, name string, serviceType enums.ServiceType) *Service {
	result := &Service{
		Environment: enums.GetEnvironment(env),
		Name:        name,
		Type:        serviceType,
	}

	return result
}

// Register is used to register an application with the router service
func (s *Service) Register(port string) error {
	err := s.setURL(port)

	if err != nil {
		return err
	}

	contents, err := POSTMessage(s.ID, "Router.API", "discovery", s)

	if err != nil {
		return err
	}

	data := MarshalToResult(contents)

	if data.Failed() {
		return data
	}

	s.ID = data.Data.(string)

	return nil
}

func (s *Service) setURL(port string) error {
	url, err := getPublicIP(port, s.Environment)

	if err != nil {
		return err
	}

	s.URL = url

	return nil
}

func getPublicIP(port string, env enums.Environment) (string, error) {
	if env == enums.DEV {
		return makeURL("localhost", port), nil
	}

	resp, err := http.Get("http://myexternalip.com/raw")

	if err != nil {
		return "error", err
	}

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "error", err
	}

	result := strings.Replace(string(ip), "\n", "", -1)

	return makeURL(result, port), nil
}

func makeURL(domain, port string) string {
	schema := "https"

	if domain == "localhost" {
		schema = "http"
	}

	return fmt.Sprintf("%s://%s:%s/", schema, domain, port)
}
