package mango

import (
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

	url, err := GetServiceURL(instanceID, "Theme.API", false)

	if err != nil {
		return err
	}

	for _, v := range lst {
		err = downloadTemplate(instanceID, v, url)

		if err != nil {
			return err
		}
	}

	return nil
}

func findTemplates(instanceID string) ([]string, error) {
	result := []string{}
	fail, err := DoGET(&result, instanceID, "Theme.API", "asset", "html")

	if err != nil {
		return result, err
	}

	if fail != nil {
		return result, fail
	}

	return result, nil
}

func downloadTemplate(instanceID, template, themeURL string) error {
	fullURL := fmt.Sprintf("%sv1/%s/%s/%s", themeURL, "asset", "html", template)
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
