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
	// 插件
	Plugins []string
	// 后置插件
	After []string

	Method string

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

func GetAllRouter() map[string]*Router {
	return m
}

func Register(serverName, endpoint, path, method string, plugins []string, req interface{}, resp interface{}, after ...string) {
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
		respType = reflect.TypeOf(resp)
		if respType.Kind() == reflect.Ptr {
			respType = respType.Elem()
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
		Method:     method,
		Plugins:    make([]string, len(plugins)),
		After:      make([]string, len(after)),
	}
	copy(r.Plugins, plugins)
	copy(r.After, after)
	m[path] = r
}

func RouterServer(key string) (Router, error) {
	if val, ok := m[key]; ok {
		return *val, nil
	}
	return Router{}, keyreflect.NoReflectError
}
