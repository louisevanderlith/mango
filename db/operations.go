package db

import (
	"github.com/astaxie/beego/orm"
	"reflect"
	"github.com/astaxie/beego"
	"strings"
	"time"
	"log"
)

var (
	persistData bool
)

func init() {
	persistData = beego.BConfig.RunMode != "dev"
}

func insert(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(obj)
}

func read(obj interface{}, related ...string) error {
	readColumns := getReadColumns(obj)

	o := orm.NewOrm()
	err := o.Read(obj, readColumns...)

	for _, v := range related {
		o.LoadRelated(obj, v)
	}

	return err
}

func readAll(filter interface{}, container interface{}) error {
	readColumns := getFilterValues(filter)

	o := orm.NewOrm()
	qt := o.QueryTable(filter)

	for k, v := range readColumns {
		qt = qt.Filter(strings.ToLower(k), v)
	}

	qt = qt.Filter("deleted", false)
	_, err := qt.All(container)

	return err
}

func update(obj interface{}) (int64, error) {
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

func isFieldSet(field interface{}) (result bool) {
	val := reflect.ValueOf(field)

	switch val.Kind() {
	case reflect.Int:
		iField := field.(int)
		result = intRule(iField)
	case reflect.Int64:
		i64Field := field.(int64)
		result = int64Rule(i64Field)
	case reflect.String:
		strField := field.(string)
		result = strRule(strField)
	case reflect.Ptr, reflect.Struct, reflect.Slice:
		result = true
	case reflect.Bool:
		result = boolRule(field)
	default:
		result = nilRule(field)
	}

	if tField, ok := field.(time.Time); ok {
		result = tField.IsZero()
	}

	return !result
}

func nilRule(val interface{}) bool {
	return val == nil
}

func strRule(val interface{}) bool {
	return val == ""
}

func intRule(val int) bool {
	return val < 1
}

func int64Rule(val int64) bool {
	return val < 1
}

func boolRule(val interface{}) bool {
	return val == false
}

func addToMap(target map[string]interface{}, items map[string]interface{}) {
	for k, v := range items {
		target[k] = v
	}
}
