package reflect

import (
	// "ipopcloud/common/utils"
	"reflect"
)
func ExtractUpdates(arg interface{}, fields []string) map[string]interface{} {
	result := make(map[string]interface{})
	r_value := reflect.Indirect(reflect.ValueOf(arg))
	r_type := r_value.Type()
	if r_type.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < r_type.NumField(); i++ {
		field := r_type.Field(i)
		name := r_type.Field(i).Name
		jsonTagName := field.Tag.Get("json")

		if !Contains(jsonTagName, fields) {
			continue
		}
		if !r_value.FieldByName(name).IsValid() {
			continue
		}

		fieldValue := r_value.FieldByName(name)
		_kind := fieldValue.Type().Kind()
		if _kind == reflect.Ptr {
			result[field.Tag.Get("json")] = r_value.FieldByName(name).Elem().Interface()
		} else {

			result[field.Tag.Get("json")] = r_value.FieldByName(name).Interface()
		}
	}
	return result
}
func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}