// Package structs provides various functions for structs.
package structs

import "reflect"

// ExtractEmbedValue extracts the value of embed type from a struct.
//
// it takes `str` as a struct with embed type and returns the extracted value `val` of generic type `T` and a boolean `ok` indicating whether the extraction was successful.
func ExtractEmbedValue[T any](str any) (val T, ok bool) {
	valueStr := reflect.ValueOf(str)
	typeOfStr := valueStr.Type()

	if typeOfStr.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < valueStr.NumField(); i++ {
		field := typeOfStr.Field(i)

		if field.Name == reflect.ValueOf(val).Type().Name() && field.Anonymous {
			fieldValue := valueStr.Field(i).Interface()
			val, ok = fieldValue.(T)

			return
		}

		if field.Type.Kind() == reflect.Struct {
			if val, ok = ExtractEmbedValue[T](valueStr.Field(i).Interface()); ok {
				return
			}
		}
	}

	return
}
