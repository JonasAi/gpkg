package utils

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
)

// 判断空值
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Array:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0+0i
	case reflect.Interface, reflect.Ptr, reflect.Chan, reflect.Map, reflect.Slice:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// 对象转 interface切片
func ToInterfaceArr(arr interface{}) []interface{} {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return nil
	}

	arrValue := reflect.ValueOf(arr)
	retArr := make([]interface{}, arrValue.Len())
	for k := 0; k < arrValue.Len(); k++ {
		retArr[k] = arrValue.Index(k).Interface()
	}
	return retArr
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
