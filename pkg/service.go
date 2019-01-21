package mango

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/louisevanderlith/mango/enums"
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

	resp, err := sendRegistration(s)

	if err != nil {
		panic(err)
	}

	if resp.Failed() {
		return resp
	}

	s.ID = resp.Data.(string)

	return nil
}

func sendRegistration(s *Service) (*RESTResult, error) {
	bits, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	routrURL, err := GetServiceURL(s.ID, "Router.API", false)

	if err != nil {
		return nil, err
	}

	disco := fmt.Sprintf("%sv1/discovery/", routrURL)
	resp, err := http.Post(disco, "application/json", bytes.NewBuffer(bits))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	data, err := MarshalToResult(contents)

	return data, err
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
