package mango

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//RESTResult is the base object of every response.
type RESTResult struct {
	Code   int         `json:"Code"`
	Reason error       `json:"Error"`
	Data   interface{} `json:"Data"`
}

func NewRESTResult(code int, reason error, data interface{}) *RESTResult {
	result := &RESTResult{
		Code: code,
		Data: data,
	}

	if reason != nil {
		result.Reason = reason
	}

	return result
}

//Failed will return true if it's found a reason.
/*func (r *RESTResult) Failed() bool {
	hasReason := len(r.Reason) > 0
	over
}*/

//DoGET does a GET request and will update the container with the reponse's values.
//returns int : httpStatusCode
//return error: error
func DoGET(container interface{}, instanceID, serviceName, controller string, params ...string) (int, error) {
	url, err := GetServiceURL(instanceID, serviceName, false)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

	resp, err := http.Get(fullURL)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	rest, err := marshalToResult(contents, container)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return rest.Code, rest.Reason
}

func marshalToResult(content []byte, dataObj interface{}) (*RESTResult, error) {
	result := &RESTResult{Data: dataObj}
	err := json.Unmarshal(content, result)

	if err != nil {
		fullerr := fmt.Errorf("marshal: %s\r%s", err.Error(), string(content))
		return nil, fullerr
	}

	return result, nil
}
