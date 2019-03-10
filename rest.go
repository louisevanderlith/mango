package mango

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

func NewRESTValue(container interface{}) *RESTResult {
	return &RESTResult{Data: container}
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

//DoGET does a GET request and will update the container with the reponse's values.
func DoGET(container interface{}, instanceID, serviceName, controller string, params ...string) (apiErr error, err error) {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return nil, err
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

	rest, err := marshalToResult(contents, container)

	if err != nil {
		return nil, err
	}

	if rest.Failed() {
		return rest, nil
	}

	return nil, nil
}

func marshalToResult(content []byte, dataObj interface{}) (*RESTResult, error) {
	result := NewRESTValue(dataObj)
	err := json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RESTResult) Error() string {
	return r.Reason
}
