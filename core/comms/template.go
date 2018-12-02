package comms

import (
	"io/ioutil"
)

func populatTemplate(msg Message) string {
	template, _ := rawTemplate()
	render := template

	return render
}

func rawTemplate() (string, error) {
	result := ""
	file, err := ioutil.ReadFile("general.html")

	if err == nil {
		result = string(file)
	}

	return result, err
}
