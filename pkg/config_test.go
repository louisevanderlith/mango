package util

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	input := "123"
	expect := 123

	actual := parse(input)

	if actual != expect {
		t.Errorf("Expected:%v Got:%v", expect, actual)
	}
}

func TestGetFileContent_MustGetLines(t *testing.T) {
	content := GetFileContent("../api/proxy/app.conf")

	if len(content) <= 0 {
		t.Error("No Config file read")
	}
}

func TestGetFileContent_MustHaveValue(t *testing.T) {
	contentbytes := GetFileContent("../api/proxy/app.conf")
	content := strings.Split(string(contentbytes), "\r\n")

	for _, val := range content {
		t.Log(val)

		if val == "" {
			t.Error("Line was empty")
		}
	}
}

func TestConfigLoad(t *testing.T) {
	config := new(Config)
	config.LoadConfig("../api/proxy/app.conf")

	t.Log(config)

	if config.Host == "" {
		t.Error("Expecting Host value")
	}
}
