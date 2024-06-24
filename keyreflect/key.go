// Package reflect
// @Author: yinwei
// @File: key
// @Version: 1.0.0
// @Date: 2024/6/24 18:37

package keyreflect

import "reflect"

var m = make(map[string]reflect.Type)

func Register(key string, t reflect.Type) {
	m[key] = t
}
