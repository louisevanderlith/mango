package secure

import (
	"testing"
)

func TestCreateUser_MustHaveError(t *testing.T) {
	input := User{
		Email:    "",
		Name:     "",
		Password: "pass"}

	err := CreateUser(input)

	if err == nil {
		t.Error("Expecting Email, Contact Number or Password invalid.")
	}
}

func TestCreateUser_PasswordLength_MustHaveError(t *testing.T) {
	input := User{
		Email:    "testing@mail.com",
		Password: "short",
	}

	input := NewUser("Jan", "testing@mail.com")
	err := NewUser(input)

	if err == nil {
		t.Error("Expecting 'Password must be atleast 6 characters'.")
	}
}

func TestSecurePassword_NoError(t *testing.T) {
	input := User{
		Password: "password"}

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
		Password:      pwd}

	CreateUser(input)

	_, err := Login(input.Email, []byte(input.Password), "", "")

	if err != nil {
		t.Error(err)
	}

	dropUser(input)
}

func TestLoginUser_Contact_NoError(t *testing.T) {
	pwd := "longPassword"
	input := User{
		ContactNumber: "0123456789",
		Email:         "test@mail.com",
		Password:      pwd}

	CreateUser(input)

	_, err := Login(input.ContactNumber, []byte(input.Password), "", "")

	if err != nil {
		t.Error(err)
	}

	dropUser(input)
}
