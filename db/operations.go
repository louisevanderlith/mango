package db

import (
	"github.com/astaxie/beego/orm"
	"reflect"
	"github.com/astaxie/beego"
	"strings"
)

var (
	rules       []func(interface{}) bool
	persistData bool
)

func init() {
	rules = getRules()
	persistData = beego.BConfig.RunMode != "dev"
}

func insert(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(obj)
}

func read(obj interface{}) error {
	readColumns := getReadColumns(obj)

	o := orm.NewOrm()

	err := o.Read(obj, readColumns...)

	return err
}

func readAll(filter interface{}, container interface{}) error {
	readColumns := getFilterValues(filter)

	o := orm.NewOrm()
	tableName := getTableName(filter)
	qt := o.QueryTable(tableName)

	for k, v := range readColumns {
		qt.Filter(k, v)
	}

	qt.Filter("deleted", false)
	_, err := qt.All(&container)

	return err
}

func update(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.Update(obj)
}

func getTableName(obj interface{}) string {
	t := reflect.TypeOf(obj)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return strings.ToLower(t.Name())
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

				if field.Name == "ID" {
					break
				}
			}
		}
	}

	return result
}

func getFilterValues(filter interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	valOf := reflect.ValueOf(filter)
	indi := reflect.Indirect(valOf)
	typeOf := indi.Type()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		if field.Anonymous {
			anon := reflect.New(field.Type).Elem().Interface()
			names := getFilterValues(anon)

			addToMap(result, names)
		} else {
			fieldVal := indi.Field(i).Interface()

			if isFieldSet(fieldVal) {
				result[field.Name] = fieldVal
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

func addToMap(target map[string]interface{}, items map[string]interface{}) {
	for k, v := range items {
		target[k] = v
	}
}
