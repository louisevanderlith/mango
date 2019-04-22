package mango

import (
	"fmt"
	"io"
	"net/http"
	"os"

	folio "github.com/louisevanderlith/folio/core"
)

func GetDefaultTheme(instanceID, siteName string) (ThemeSetting, error) {
	prof := folio.Profile{}
	err := DoGET(&prof, instanceID, "Folio.API", "profile", siteName)

	if err != nil {
		return ThemeSetting{}, err
	}

	result := ThemeSetting{
		InstanceID: instanceID,
		LogoKey:    prof.ImageKey,
		Name:       prof.Title,
	}

	return result, nil
}

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
	err := DoGET(&result, instanceID, "Theme.API", "asset", "html")

	if err != nil {
		return result, err
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
