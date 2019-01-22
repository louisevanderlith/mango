package mango

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
