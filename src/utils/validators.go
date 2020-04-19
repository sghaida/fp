package utils

import (
	"reflect"
)

// IsNilOrEmpty checks if data is empty or nil
func IsNilOrEmpty(object interface{}) bool {
	if object == nil {
		return true
	}
	if object == "" {
		return true
	}
	// extract the value of the interface
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		return reflect.DeepEqual(object, empty)
	}
	switch reflect.TypeOf(object).Kind() {
	case reflect.Ptr, reflect.Interface:
		return IsNilOrEmpty(reflect.ValueOf(object).Elem())
	case reflect.Slice, reflect.Array, reflect.Map:
		return reflect.ValueOf(object).IsNil() || reflect.ValueOf(object).Len() == 0
	}
	return false
}
