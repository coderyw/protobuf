// Package log
// @Author: yinwei
// @File: logger
// @Version: 1.0.0
// @Date: 2024/7/15 13:06

package log

type Logger interface {
	Error(msg string, err error)
	Info(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

var logger Logger

func WithLogger(l Logger) {
	logger = l
}

func Error(msg string, err error) {
	if logger == nil {
		return
	}
	logger.Error(msg, err)
}

func Info(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Info(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Fatalf(msg, args...)
}
