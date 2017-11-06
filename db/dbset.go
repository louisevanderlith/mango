package db

import (
	"reflect"

	"github.com/astaxie/beego/orm"
)

type DBSet interface {
	Insert() (int64, error)
	Read() error
	ReadlAll(*[]interface, error)
	Update() (int64, error)
	Delete() error
}

var rules []func(interface{}) bool

func init() {
	rules = getRules()
}

func Insert(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(obj)
}

func Read(obj interface{}) error {
	readColumns := getReadColumns(obj)

	o := orm.NewOrm()

	err := o.Read(obj, readColumns...)

	return err
}

func ReadAll(obj interface{}) (*[]interface{}, error) {
	var result []interface{}

	o := orm.NewOrm()
	qs := o.QueryTable(obj).Filter("Deleted", false)

	_, err := qs.All(&result)

	return result, err
}

func Update(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.Update(obj)
}

func getReadColumns(obj interface{}) []string {
	var result []string

	valOf := reflect.ValueOf(obj)
	indi := reflect.Indirect(valOf)
	typeOf := indi.Type()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		if field.Anonymous {
			anon := reflect.New(field.Type).Elem().Interface()
			names := getReadColumns(anon)

			result = append(result, names...)
		} else {
			fieldVal := indi.Field(i).Interface()

			if isFieldSet(fieldVal) {
				result = append(result, field.Name)
			}
		}
	}

	return result
}

func isFieldSet(field interface{}) bool {
	result := true

	for _, v := range rules {
		if v(field) {
			result = false
			break
		}
	}

	return result
}

func getRules() []func(interface{}) bool {
	rules := []func(interface{}) bool{
		nilRule,
		strRule,
		intRule,
		boolRule}

	return rules
}

func strRule(val interface{}) bool {
	return val == ""
}

func intRule(val interface{}) bool {
	return val == 0
}

func nilRule(val interface{}) bool {
	return val == nil
}

func boolRule(val interface{}) bool {
	return val == false
}
