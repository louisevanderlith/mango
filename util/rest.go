package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GETMessage(serviceName, controller string, params ...string) (result []byte, err error) {
	url, err := GetServiceURL(serviceName, false)

	if err != nil {
		return result, err
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fullURL, nil)

	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	return contents, err
}

func POSTMessage(serviceName, controller string, obj interface{}) (result []byte, err error) {
	url, err := GetServiceURL(serviceName, false)

	if err != nil {
		return []byte{}, err
	}

	fullURL := fmt.Sprintf("%sv1/%s", url, controller)

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(obj)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", fullURL, buff)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	return contents, err
}

type RESTResult struct {
	Error string
	Data  interface{}
}

func MarshalToResult(content []byte) RESTResult {
	result := RESTResult{}
	err := json.Unmarshal(content, &result)

	if err != nil {
		return RESTResult{err.Error(), nil}
	}

	return result
}
