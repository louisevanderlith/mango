package util

import (
	"bytes"
	"encoding/json"
	"errors"
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
func GETMessage(instanceID, serviceName, controller string, params ...string) *RESTResult {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

	return jsonRequest("GET", fullURL, "")
}

// POSTMessage does a POST Request
func POSTMessage(instanceID, serviceName, controller string, obj interface{}) *RESTResult {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	fullURL := fmt.Sprintf("%sv1/%s", url, controller)

	return jsonRequest("POST", fullURL, obj)
}

func jsonRequest(action, url string, obj interface{}) *RESTResult {
	buff := &bytes.Buffer{}
	req := &http.Request{}
	err := json.NewEncoder(buff).Encode(obj)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	req, err = http.NewRequest(action, url, buff)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	if action == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return NewRESTResult(err, nil)
	}

	data := MarshalToResult(contents)

	if data.Failed() {
		return NewRESTResult(data, nil)
	}

	if data.Data == nil {
		return NewRESTResult(errors.New("data.Data is nil"), nil)
	}

	return data
}

func MarshalToResult(content []byte) *RESTResult {
	result := &RESTResult{}
	err := json.Unmarshal(content, result)

	if err != nil {
		return NewRESTResult(err, result)
	}

	return result
}

func (r *RESTResult) Error() string {
	return r.Reason
}
