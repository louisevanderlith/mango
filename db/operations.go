package db

import (
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	persistData bool
)

func init() {
	persistData = beego.BConfig.RunMode != "dev"
}

func insert(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	relationships := getRelationships(obj)

	for _, v := range relationships {
		o.Insert(v)
	}

	return o.Insert(obj)
}

func insertMulti(batchCount int, objs interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.InsertMulti(batchCount, objs)
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
	relationships := getRelationships(obj)

	for _, v := range relationships {
		o.Update(v, getReadColumns(v)...)
	}

	return o.Update(obj, getReadColumns(obj)...)
}

// getRelationships returns the objects related to the current object
func getRelationships(obj interface{}) []interface{} {
	var result []interface{}

	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Type.Kind() == reflect.Ptr || typeField.Type.Kind() == reflect.Slice {
			value := valueField.Interface()
			if value != nil {
				switch reflect.TypeOf(value).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(value)

					for j := 0; j < s.Len(); j++ {
						result = append(result, s.Index(j).Interface())
					}
				default:
					result = append(result, value)
				}
			}
		}
	}

	return result
}

// getReadColumns returns a list of column names, used to search
func getReadColumns(obj interface{}) []string {
	var result []string

	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := valueField.Interface()
		kind := valueField.Kind()

		if !isFieldEmpty(value, kind) {
			result = append(result, typeField.Name)
		}
	}

	return result
}

// getFilterValues returns a map of parameters and their values, used to filter.
func getFilterValues(filter interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(filter).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := valueField.Interface()
		kind := valueField.Kind()

		if !isFieldEmpty(value, kind) {
			result[typeField.Name] = value
		}
	}

	return result
}

/*
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

			if !isFieldEmpty(fieldVal) {
				result[field.Name] = fieldVal
			}
		}
	}

	return result
}*/

/*
func getReadColumnsOld(obj interface{}) []string {
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
		} else if field.PkgPath != "reflect" {
			fmt.Println(field.PkgPath)
			indiField := indi.Field(i)
			fieldVal := indiField.Interface()

			if isFieldSet(fieldVal) {
				result = append(result, field.Name)

				if field.Name == "ID" {
					break
				}
			}
		}
	}

	return result
}*/

func isFieldEmpty(val interface{}, kind reflect.Kind) bool {
	var result bool

	switch kind {
	case reflect.Int:
		iField := val.(int)
		result = intRule(iField)
	case reflect.Int64:
		i64Field := val.(int64)
		result = int64Rule(i64Field)
	case reflect.String:
		strField := val.(string)
		result = strRule(strField)
	case reflect.Ptr, reflect.Struct, reflect.Slice:
		result = true
	case reflect.Bool:
		result = boolRule(val)
	default:
		result = nilRule(val)
	}

	if tField, ok := val.(time.Time); ok {
		result = tField.IsZero()
	}

	return result
}

func sliceRule(val interface{}) bool {
	records := reflect.ValueOf(val)
	return records.Len() <= 0
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
