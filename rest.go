package mango

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//RESTResult is the base object of every response.
type RESTResult struct {
	Code   int         `json:"Code"`
	Reason string      `json:"Error"`
	Data   interface{} `json:"Data"`
}

var client = &http.Client{}

func NewRESTResult(code int, reason error, data interface{}) *RESTResult {
	result := &RESTResult{
		Code: code,
		Data: data,
	}

	if reason != nil {
		result.Reason = reason.Error()
	}

	return result
}

func (r RESTResult) Error() string {
	return r.Reason
}

//DoGET does a GET request and will update the container with the reponse's values.
//Mango only exposes GET, as other requests should be made from the Client's Browser
//token: this is the access_token/avosession
//container: the object that will be populated with the results
//instanceID: instance of the application making the request
//serviceName: the name of the service being requested
//controller: the Controller to call
//params: additional path variables
//returns int : httpStatusCode
//return error: error
func DoGET(token string, container interface{}, instanceID, serviceName, controller string, params ...string) (int, error) {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, errors.New(string(contents))
	}

	rest, err := marshalToResult(contents, container)

	if err != nil {
		msg := fmt.Errorf("Invalid JSON; Body:\n%s\nError:\n%s", string(contents), err)
		return http.StatusInternalServerError, msg
	}

	if len(rest.Reason) > 0 {
		return rest.Code, rest
	}

	return rest.Code, nil
}

func marshalToResult(content []byte, dataObj interface{}) (*RESTResult, error) {
	result := &RESTResult{Data: dataObj}
	err := json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
