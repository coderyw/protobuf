// Package route
// @Author: yinwei
// @File: path
// @Version: 1.0.0
// @Date: 2024/7/17 09:29

package route

import (
	"github.com/coderyw/protobuf/keyreflect"
	"reflect"
)

// 路由对应type

var m = make(map[string]*Router)
var cache = make(map[string]reflect.Type)

type Router struct {
	// server name
	ServerName string
	// 端点
	Endpoint string
	// 路由path
	Path string

	req  reflect.Type
	resp reflect.Type
}

func (r *Router) ReqI() (interface{}, error) {
	if r.req == nil {
		return nil, keyreflect.NoReflectError
	}
	return reflect.New(r.req).Interface(), nil
}
func (r *Router) RespI() (interface{}, error) {
	if r.resp == nil {
		return nil, keyreflect.NoReflectError
	}
	return reflect.New(r.resp).Interface(), nil
}

func Register(serverName, endpoint, path string, req interface{}, resp interface{}) {
	var reqType reflect.Type
	var respType reflect.Type
	if req != nil {
		reqType = reflect.TypeOf(req)
		if reqType.Kind() == reflect.Ptr {
			reqType = reqType.Elem()
		}
		if val, ok := cache[reqType.String()]; ok {
			reqType = val
		} else {
			cache[reqType.String()] = reqType
		}
	}
	if resp != nil {
		respType = reflect.TypeOf(respType)
		if respType.Kind() == reflect.Ptr {
			respType = reqType.Elem()
		}
		if val, ok := cache[respType.String()]; ok {
			respType = val
		} else {
			cache[respType.String()] = respType
		}
	}

	r := &Router{
		ServerName: serverName,
		Endpoint:   endpoint,
		Path:       path,
		req:        reqType,
		resp:       respType,
	}

	m[path] = r
}

func RouterServer(key string) (Router, error) {
	if val, ok := m[key]; ok {
		return *val, nil
	}
	return Router{}, keyreflect.NoReflectError
}
