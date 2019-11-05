package helper

import (
	"fmt"
	"reflect"
)

// convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

func ToInt(val interface{}) int {
	if val, ok := val.(int8); ok {
		return int(val)
	}
	if val, ok := val.(int16); ok {
		return int(val)
	}
	if val, ok := val.(int32); ok {
		return int(val)
	}
	if val, ok := val.(int64); ok {
		return int(val)
	}
	return 0
}