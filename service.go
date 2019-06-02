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

//Service identifies the Registering APP or API
type Service struct {
	ID             string
	Name           string
	URL            string
	Version        int
	AllowedCallers map[enums.ServiceType]struct{}
	Type           enums.ServiceType
	PublicKey      string
}

//NewService returns a new instance of a Services' information
//publicKey refers to the location of the public key file (.pub)
func NewService(name, publicKey string, serviceType enums.ServiceType) *Service {
	result := &Service{
		Name:           name,
		Type:           serviceType,
		PublicKey:      publicKey,
		AllowedCallers: make(map[enums.ServiceType]struct{}),
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
		return err
	}

	if resp.Code != http.StatusOK && len(resp.Reason) != 0 {
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

	data, err := marshalToResult(contents, "")

	return data, err
}

func (s *Service) setURL(port string) error {
	url, err := getNetworkIP(s.Name, port)

	if err != nil {
		return err
	}

	s.URL = url

	return nil
}

func getNetworkIP(name, port string) (string, error) {
	uniqueName := strings.Replace(name, ".", "", 1)

	return makeURL(uniqueName, port), nil
}

func makeURL(domain, port string) string {
	schema := "http"

	return fmt.Sprintf("%s://%s:%s/", schema, domain, port)
}
