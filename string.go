package zutil

import (
	"math/rand"
	"reflect"
	"strconv"
)

<<<<<<< HEAD
=======
// MapToString 参数识别
>>>>>>> 8bb048339a0bdf55f74c4da1769717df9866f515
func MapToString(arguments map[string]interface{}) string {
	str := ""
	for k, v := range arguments {
		s := ""
		v := reflect.ValueOf(v)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			s += strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			s += strconv.FormatUint(v.Uint(), 10)
		case reflect.Bool:
			s += strconv.FormatBool(v.Bool())
		case reflect.String:
			s += strconv.Quote(v.String())
		case reflect.Float32, reflect.Float64:
			s += strconv.Itoa(int(v.Float()))
		}
		if len(s) > 0 {
			str += " -" + k + " " + s
		}
	}
	return str
}
<<<<<<< HEAD
func MapToStringV2(arguments map[string]interface{}) string {
	str := ""
	for k, v := range arguments {
		s := ""
		v := reflect.ValueOf(v)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			s += strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			s += strconv.FormatUint(v.Uint(), 10)
		case reflect.Bool:
			s += strconv.FormatBool(v.Bool())
		case reflect.String:
			s += strconv.Quote(v.String())
		case reflect.Float32, reflect.Float64:
			s += strconv.Itoa(int(v.Float()))
		}
		if len(s) > 0 {
			if len(k) > 1 {
				str += " --" + k + " " + s
			} else {
				str += " -" + k + " " + s
			}
		}
	}
	return str
=======

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// ContainsString 返回list中是否包含
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
>>>>>>> 8bb048339a0bdf55f74c4da1769717df9866f515
}
