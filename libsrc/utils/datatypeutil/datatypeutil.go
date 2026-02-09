package datatypeutil

import "reflect"

func IsString(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.String {
		return true
	}
	return false
}

func IsInt(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Int || t.Kind() == reflect.Int64 {
		return true
	}
	return false
}

func IsFloat(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64 {
		return true
	}
	return false
}

func IsObject(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Map {
		return true
	}
	return false
}

func IsObjectArray(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Slice {
		return true
	}
	return false
}

func IsBool(data interface{}) bool {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Bool {
		return true
	}
	return false
}
