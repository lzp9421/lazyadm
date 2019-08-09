package library

import (
	"strconv"
	"fmt"
	"reflect"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func  ToString(arg interface{}) string {
	switch arg.(type) {
	case int:
		return strconv.Itoa(arg.(int))
	case int64:
		return strconv.FormatInt(arg.(int64), 10)
	case []byte:
		return string(arg.([]byte))
	case string:
		return arg.(string)
	default:
		return ""
	}
}

func Addslashes(v string) string {
	pos := 0
	buf := make([]byte, len(v)*2)
	for i := 0; i < len(v); i++ {
		c := v[i]
		if c == '\'' || c == '"' || c == '\\' {
			buf[pos] = '\\'
			buf[pos+1] = c
			pos += 2
		} else {
			buf[pos] = c
			pos++
		}
	}
	return string(buf[:pos])
}


func ToInt64(value interface{}) int64 {
	val := reflect.ValueOf(value)
	var d int64
	var err error
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	if err != nil {
		d = 0
	}
	return d
}

func ToInt(value interface{}) int {
	val := reflect.ValueOf(value)
	var d int
	var err error
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = int(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		d = int(val.Uint())
	case string:
		d, err = strconv.Atoi(val.String())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	if err != nil {
		d = 0
	}

	return d
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func ToSnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func ToCamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func ContainInArray(str string, arr []string) bool {
	for _, s := range arr  {
		if strings.Index(s, str) != -1 {
			return true
		}
	}
	return false
}

func InArray(str string, arr []string) bool {
	for _, s := range arr  {
		if s == str {
			return true
		}
	}
	return false
}
