package util

import (
	"encoding/json"
	"log"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"bytes"
)

func GETMessage(serviceName, controller string, params ...string) ([]byte, int) {
	var result []byte
	var statusCode int
	url, err := GetServiceURL(serviceName)

	if err == nil {
		fullURL := fmt.Sprintf("%sv1/%s/%s", url, controller, strings.Join(params, "/"))

		client := &http.Client{}
		req, _ := http.NewRequest("GET", fullURL, nil)

		resp, err := client.Do(req)

		if err != nil {
			log.Printf("http.Get: ", err)
		} else {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Printf("ioutil.ReadAll: ", err)
			}

			result = contents
		}
	} else {
		statusCode = 500
		result = []byte(err.Error())
	}

	return result, statusCode
}

func POSTMessage(serviceName, action string, obj interface{}) ([]byte, int) {
	var result []byte
	var statusCode int
	url, err := GetServiceURL(serviceName)

	if err == nil {
		fullURL := fmt.Sprintf("%sv1/%s", url, action)

		buff := new(bytes.Buffer)
		json.NewEncoder(buff).Encode(obj)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", fullURL, buff)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)

		if err != nil {
			log.Printf("http.Post: ", err)
		} else {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Printf("ioutil.ReadAll: ", err)
			}

			result = contents
		}
	} else {
		statusCode = 500
		result = []byte(err.Error())
	}

	return result, statusCode
}

func MarshalToMap(content []byte) map[string]*json.RawMessage {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(content, &objmap)

	if err != nil {
		log.Printf("MarshalToMap: ", err)
	}

	return objmap
}
