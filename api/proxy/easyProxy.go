package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Target   string
	Data     string
	OriginID string
	Action   string
}

type ResultMessage struct {
	Success   bool
	Data      string
	Error     string
	RequestID string
}

func Proxer() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(config.Host, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	response := ResultMessage{
		Success: true}

	message, err := getMessage(r.Body)

	if err != nil {
		response.Error = err.Error()
		response.Success = false
	} else {
		data, reqErr := forwardRequest(message)

		if reqErr != nil {
			response.Error = reqErr.Error()
			response.Success = false
		}

		response.Data = data
	}

	json.NewEncoder(w).Encode(response)
}

func getMessage(body io.ReadCloser) (Message, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)

	var message Message
	err := decoder.Decode(&message)

	if err != nil {
		log.Fatal(err)
	}

	return message, err
}

func forwardRequest(message Message) (string, error) {
	var data string
	var err error

	target, err := getTargetURL(message.OriginID, message.Target)

	if err == nil {
		if message.Action == "GET" {
			data, err = forwardGET(target, message)
		} else if message.Action == "POST" {
			data, err = forwardPOST(target, message)
		} else {
			msg := fmt.Sprintf("Action of %s not supported", message.Action)
			err = errors.New(msg)
		}
	}

	return data, err
}

func forwardPOST(target string, message Message) (string, error) {
	var result string
	var err error

	reader := strings.NewReader(message.Data)
	resp, err := http.Post(target, "application/json", reader)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	} else {

		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		result = string(contents)
	}

	return result, err
}

func forwardGET(target string, message Message) (string, error) {
	var result string
	var err error

	resp, err := http.Get(target)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	} else {

		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		result = string(contents)
	}

	return result, err
}

func getTargetURL(appKey string, rawTarget string) (string, error) {
	var result string
	var err error

	discoveryRoute := fmt.Sprintf("%s%s/%s", config.Discovery, appKey, rawTarget)
	resp, err := http.Get(discoveryRoute)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		result = string(contents)

		if result == "" {
			msg := fmt.Sprintf("Couldn't find a application for %s", rawTarget)
			err = errors.New(msg)
		}
	}

	return result, err
}
