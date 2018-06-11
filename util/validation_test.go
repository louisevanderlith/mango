package util

import "testing"

func TestValidateStruct_StringValid(t *testing.T) {
	obj := struct {
		Name string `orm:"null;size(50)"`
	}{
		Name: "ABC",
	}

	valid, err := ValidateStruct(obj)

	if !valid {
		t.Error("Expecting obj to be valid.", err)
	}
}

func TestValidateStruct_StringRequired(t *testing.T) {
	obj := struct {
		Name string `orm:"size(50)"`
	}{
		Name: "",
	}

	valid, _ := ValidateStruct(obj)

	if valid {
		t.Error("Expecting error 'Name is required'")
	}
}

func TestValidateStruct_StringTooLong(t *testing.T) {
	obj := struct {
		Name string `orm:"size(15)"`
	}{
		Name: "1234567890ABCDEFGHIJKL",
	}

	valid, _ := ValidateStruct(obj)

	if valid {
		t.Error("Expecting error 'Name can't be more than 15 characters.'")
	}
}

func TestValidateStruct_SimpleObject_IntValid(t *testing.T) {
	obj := struct {
		Age int `orm:"null"`
	}{
		Age: 24,
	}

	valid, err := ValidateStruct(obj)

	if !valid {
		t.Error("Expecting obj to be valid.", err)
	}
}

func TestValidateStruct_SimpleObject_IntRequired(t *testing.T) {
	obj := struct {
		Age int
	}{
		Age: 0,
	}

	valid, _ := ValidateStruct(obj)

	if valid {
		t.Error("Expecting error 'Age is required'")
	}
}

func TestValidateStruct_ChildStruct_ChildValid(t *testing.T) {
	type child struct {
		Age int `orm:"null"`
	}
	obj := struct {
		Name  string `orm:"null"`
		Child *child
	}{
		Name:  "ABC",
		Child: &child{Age: 24},
	}

	valid, err := ValidateStruct(obj)

	if !valid {
		t.Error("Expecting obj to be valid", err)
	}
}

func TestValidateStruct_ChildStruct_EmptyChildValid(t *testing.T) {
	type child struct {
		Age int `orm:"null"`
	}
	obj := struct {
		Name  string `orm:"null"`
		Child *child
	}{
		Name: "ABC",
	}

	valid, err := ValidateStruct(obj)

	if !valid {
		t.Error("Expecting obj to be valid", err)
	}
}
