package util

import (
	"fmt"
	"strings"
	"errors"
	"reflect"
	"strconv"
)

type tagMeta struct {
	Required bool
	Size     int
	Type     reflect.Kind
	PropName string
}

const idMessage = "%s must be provided."
const emptyMessage = "%s can't be empty."
const shortMessage = "%s can't be more than %v characters."
const relationMessage = "%s can't be nil."
const incorrectType = "%s's value '%s' is not of type %s."

func getIDMessage(property string) string {
	return fmt.Sprintf(idMessage, property)
}

func getEmptyMessage(property string) string {
	return fmt.Sprintf(emptyMessage, property)
}

func getShortMessage(property string, length int) string {
	return fmt.Sprintf(shortMessage, property, length)
}

func getRelationMessage(property string) string {
	return fmt.Sprintf(relationMessage, property)
}

func getIncorrectType(property string, value interface{}, correctType string) string {
	return fmt.Sprintf(incorrectType, property, value, correctType)
}

func getMeta(tag string, property reflect.Value) tagMeta {
	result := tagMeta{}
	parts := strings.Split(tag, ";")

	required := !strings.Contains(tag, "null")
	result.Required = required
	result.Type = property.Type().Kind()

	hasSize := strings.Contains(tag, "size")

	if hasSize {
		rawSize := getFromTag(parts, "size")
		sSize := strings.Replace(strings.Replace(rawSize, "size(", "", -1), ")", "", -1)

		size, err := strconv.ParseInt(sSize, 10, 32)

		if err == nil {
			result.Size = int(size)
		}
	}

	return result
}

func getFromTag(list []string, name string) string {
	var result string
	for _, v := range list {
		if strings.Contains(v, name) {
			result = v
			break
		}
	}

	return result
}

func ValidateStruct(obj interface{}) (bool, error) {
	var issues []string

	v := reflect.ValueOf(obj)

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag.Get("orm")

		meta := getMeta(tag, v.Field(i))
		meta.PropName = typeField.Name

		value := v.Field(i).Interface()
		validator := getTypeValidator(meta.Type)
		isValid, problems := validator.Valid(value, meta)

		if !isValid {
			issues = append(issues, problems...)
		}
	}

	isValid := len(issues) < 1
	finErr := errors.New(strings.Join(issues, "\r\n"))

	return isValid, finErr
}

type IValidation interface {
	Valid(obj interface{}, meta tagMeta) (bool, []string)
}

func getTypeValidator(fieldType reflect.Kind) IValidation {
	var result IValidation

	switch fieldType {
	case reflect.Int:
		result = IntValidation{}
	case reflect.Int64:
		result = Int64Validation{}
	case reflect.String:
		result = StringValidation{}
	case reflect.Struct:
		result = StructValidation{}
	case reflect.Ptr:
		result = PointerValidation{}
	default:
		fmt.Println(fieldType)
		result = PointerValidation{}
	}

	return result
}

type StringValidation struct{}
type IntValidation struct{}
type Int64Validation struct{}
type StructValidation struct{}
type PointerValidation struct{}

func (o StringValidation) Valid(obj interface{}, meta tagMeta) (bool, []string) {
	var issues []string
	val, ok := obj.(string)

	if ok {
		if meta.Required && val == "" {
			issues = append(issues, getEmptyMessage(meta.PropName))
		}

		if meta.Size > 0 && len(val) > meta.Size {
			issues = append(issues, getShortMessage(meta.PropName, meta.Size))
		}
	} else {
		issues = append(issues, getIncorrectType(meta.PropName, obj, "string"))
	}

	isValid := len(issues) < 1

	return isValid, issues
}

func (o IntValidation) Valid(obj interface{}, meta tagMeta) (bool, []string) {
	var issues []string
	val, ok := obj.(int)

	if ok {
		if meta.Required && val < 1 {
			issues = append(issues, getIDMessage(meta.PropName))
		}
	} else {
		issues = append(issues, getIncorrectType(meta.PropName, obj, "int"))
	}

	isValid := len(issues) < 1

	return isValid, issues
}

func (o Int64Validation) Valid(obj interface{}, meta tagMeta) (bool, []string) {
	var issues []string
	val, ok := obj.(int64)

	if ok {
		if meta.Required && val < 1 {
			issues = append(issues, getIDMessage(meta.PropName))
		}
	} else {
		issues = append(issues, getIncorrectType(meta.PropName, obj, "int64"))
	}

	isValid := len(issues) < 1

	return isValid, issues
}

func (o StructValidation) Valid(obj interface{}, meta tagMeta) (bool, []string) {
	var issues []string

	if meta.Required && obj == nil {
		issues = append(issues, getRelationMessage(meta.PropName))
	}

	isValid := len(issues) < 1

	return isValid, issues
}

func (o PointerValidation) Valid(obj interface{}, meta tagMeta) (bool, []string) {
	return true, nil
}
