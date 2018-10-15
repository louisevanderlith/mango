package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RESTResult struct {
	Reason string      `json:"Error"`
	Data   interface{} `json:"Data"`
}

func NewRESTResult(reason error, data interface{}) *RESTResult {
	result := &RESTResult{}

	if reason != nil {
		result.Reason = reason.Error()
	}

	if data == nil {
		data = "Nothing bad happened..."
	}

	result.Data = data

	return result
}

func (r *RESTResult) Failed() bool {
	return len(r.Reason) > 0
}

// GETMessage does a GET Request
func GETMessage(instanceID, serviceName, controller string, params ...string) (*RESTResult, error) {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return nil, NewRESTResult(err, nil)
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

	resp, err := http.Get(fullURL)

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

func MarshalToResult(content []byte) (*RESTResult, error) {
	result := &RESTResult{}
	err := json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RESTResult) Error() string {
	return r.Reason
}
