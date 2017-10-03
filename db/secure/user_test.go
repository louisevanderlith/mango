package secure

import (
	"testing"
)

func TestCreateUser_MustHaveError(t *testing.T) {
	input := User{
		ContactNumber: "",
		Email:         "",
		Password:      []byte("pass")}

	err := CreateUser(input)

	if err == nil {
		t.Error("Expecting Email, Contact Number or Password invalid.")
	}
}

func TestCreateUser_PasswordLength_MustHaveError(t *testing.T) {
	input := User{
		ContactNumber: "0123457894",
		Email:         "testing@mail.com",
		Password:      []byte("short")}

	err := CreateUser(input)

	if err == nil {
		t.Error("Expecting 'Password must be atleast 6 characters'.")
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

func TestLoginUser_Email_NoError(t *testing.T) {
	pwd := "longPassword"
	input := User{
		ContactNumber: "0123456789",
		Email:         "test@mail.com",
		Password:      []byte(pwd)}

	CreateUser(input)

	loginPass := Login(input.Email, input.Password)

	if !loginPass {
		t.Error("Couldn't Login")
	}

	dropUser(input)
}

func TestLoginUser_Contact_NoError(t *testing.T) {
	pwd := "longPassword"
	input := User{
		ContactNumber: "0123456789",
		Email:         "test@mail.com",
		Password:      []byte(pwd)}

	CreateUser(input)

	loginPass := Login(input.ContactNumber, input.Password)

	if !loginPass {
		t.Error("Couldn't Login")
	}

	dropUser(input)
}
