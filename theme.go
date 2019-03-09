package mango

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

//UpdateTheme downloads the latest master templates from Theme.API
func UpdateTheme(instanceID string) error {
	lst, err := findTemplates(instanceID)

	if err != nil {
		return err
	}

	for _, v := range lst {
		err = downloadTemplate(instanceID, v)

		if err != nil {
			return err
		}
	}

	return nil
}

func findTemplates(instanceID string) ([]string, error) {
	resp, err := GETMessage(instanceID, "Theme.API", "asset", "html")

	if err != nil {
		return []string{}, err
	}

	if resp.Failed() {
		return []string{}, resp
	}

	result, ok := resp.Data.([]string)

	if !ok {
		return []string{}, errors.New("Data is not []string")
	}

	return result, nil
}

func downloadTemplate(instanceID, template string) error {
	url, err := GetServiceURL(instanceID, "Theme.API", false)

	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%sv1/%s/%s/%s", url, "asset", "html", template)
	resp, err := http.Get(fullURL)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create("/views/_shared/" + template)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
