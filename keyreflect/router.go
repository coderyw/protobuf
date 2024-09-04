// Package keyreflect
// @Author: yinwei
// @File: router
// @Version: 1.0.0
// @Date: 2024/9/3 17:29

package keyreflect

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

func RegisterGinRouter(handname string, handler interface{}) error {
	return rs.register(handname, handler)
}

func GinRouterHandler(ctx context.Context, endpoint string, decode func(interface{}) error) (interface{}, error) {

	//m := route.GetAllRouter()
	//if m == nil {
	//	return nil, fmt.Errorf("miss router")
	//}
	//r,ok:= m[path]
	//if !ok{
	//	return nil, fmt.Errorf("")
	//}

	serviceName, methodName, err := serviceMethod(endpoint)
	if err != nil {
		return nil, err
	}
	rs.mu.Lock()
	service := rs.serviceMap[serviceName]
	rs.mu.Unlock()
	if service == nil {
		return nil, fmt.Errorf("unknown router service:%v", serviceName)
	}
	mtype := service.method[methodName]
	if mtype == nil {
		return nil, fmt.Errorf("unknown service %s.%s", serviceName, methodName)
	}
	return processRequest(service, mtype, decode, ctx)
}

func processRequest(service *service, mtype *methodType, decode func(interface{}) error, ctx context.Context) (interface{}, error) {
	var argv, replyv reflect.Value
	// Decode the argument value.
	argIsValue := false // if true, need to indirect before calling.
	if mtype.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(mtype.ArgType.Elem())
	} else {
		argv = reflect.New(mtype.ArgType)
		argIsValue = true
	}
	if err := decode(argv.Interface()); err != nil {
		return nil, err
	}
	if argIsValue {
		argv = argv.Elem()
	}
	replyv = reflect.New(mtype.ReplyType.Elem())

	function := mtype.method.Func
	var returnValues []reflect.Value

	returnValues = function.Call([]reflect.Value{service.rcvr, mtype.prepareContext(ctx), argv, replyv})
	if rerr := returnValues[0].Interface(); rerr != nil {
		return nil, rerr.(error)
	}
	return replyv.Interface(), nil
}

func serviceMethod(m string) (string, string, error) {
	if len(m) == 0 {
		return "", "", fmt.Errorf("malformed method name: %q", m)
	}

	// grpc method
	if m[0] == '/' {
		// [ , Foo, Bar]
		// [ , package.Foo, Bar]
		// [ , a.package.Foo, Bar]
		parts := strings.Split(m, "/")
		if len(parts) != 3 || len(parts[1]) == 0 || len(parts[2]) == 0 {
			return "", "", fmt.Errorf("malformed method name: %q", m)
		}
		service := strings.Split(parts[1], ".")
		return service[len(service)-1], parts[2], nil
	}

	// non grpc method
	parts := strings.Split(m, ".")

	// expect [Foo, Bar]
	if len(parts) != 2 {
		return "", "", fmt.Errorf("malformed method name: %q", m)
	}

	return parts[0], parts[1], nil
}
