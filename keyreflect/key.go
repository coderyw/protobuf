// Package reflect
// @Author: yinwei
// @File: key
// @Version: 1.0.0
// @Date: 2024/6/24 18:37

package keyreflect

import (
	"fmt"
	"reflect"
)

var m = make(map[string]reflect.Type)
var NoReflectError = fmt.Errorf("no reflect type in map")

func Register(key string, t reflect.Type) {
	m[key] = t
}

func New(key string) (interface{}, error) {
	if val, ok := m[key]; ok {
		return reflect.New(val).Interface(), nil
	}
	return nil, NoReflectError
}
