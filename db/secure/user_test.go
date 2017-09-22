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

func TestSecurePassword_NoError(t *testing.T) {
	input := User{
		Password: []byte("password")}

	securePassword(&input)

	if string(input.Password) == "password" || len(input.Password) <= 0 {
		t.Error("Password was not updated")
	}
}
