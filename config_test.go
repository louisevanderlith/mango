package mango

import (
	"strings"
	"testing"
)

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
