package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GETMessage(instanceKey, serviceName, controller string, params ...string) ([]byte, error) {
	url, err := GetServiceURL(instanceKey, serviceName, false)

	if err != nil {
		return []byte{}, err
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

	return jsonRequest("GET", fullURL, nil)
}

func POSTMessage(instanceKey, serviceName, controller string, obj interface{}) ([]byte, error) {
	url, err := GetServiceURL(instanceKey, serviceName, false)

	if err != nil {
		return []byte{}, err
	}

	fullURL := fmt.Sprintf("%sv1/%s", url, controller)

	return jsonRequest("POST", fullURL, obj)
}

func jsonRequest(action, url string, obj interface{}) ([]byte, error) {
	var buff *bytes.Buffer
	var req *http.Request

	err := json.NewEncoder(buff).Encode(obj)

	if err != nil {
		return []byte{}, err
	}

	req, err = http.NewRequest(action, url, buff)

	if action == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	return contents, err
}

type RESTResult struct {
	Failed bool
	Reason string
	Data   interface{}
}

func NewRESTResult(reason string, data interface{}) *RESTResult {
	result := &RESTResult{
		Failed: len(reason) > 0,
		Reason: reason,
		Data:   data,
	}

	return result
}

func MarshalToResult(content []byte) *RESTResult {
	result := &RESTResult{}
	err := json.Unmarshal(content, result)

	if err != nil {
		return NewRESTResult(err.Error(), result)
	}

	return result
}

func (r *RESTResult) Error() string {
	return r.Reason
}
