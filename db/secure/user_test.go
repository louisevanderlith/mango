package secure

import (
	"testing"
)

func TestCreateUser_MustHaveError(t *testing.T) {
	input := User{
		ContactNumber: "",
		Email:         ""}

	err := CreateUser(input)

	if err == nil {
		t.Errorf("Expecting Email or Contact Number invalid. Got: %s", err)
	}
}
