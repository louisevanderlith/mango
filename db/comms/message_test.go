package comms

import "testing"

func TestRawTemplate_NotEmpty(t *testing.T) {
	actual, err := rawTemplate()

	if err != nil {
		t.Error(err)
	}

	if actual == "" {
		t.Error("template is empty")
	}
}

func TestPopulateTemplate(t *testing.T) {
	actual, err := rawTemplate()

	if err != nil {
		t.Error(err)
	}

	if actual == "" {
		t.Error("msg is empty")
	}
}
