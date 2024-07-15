// Package log
// @Author: yinwei
// @File: logger
// @Version: 1.0.0
// @Date: 2024/7/15 13:06

package log

type Logger interface {
	Error(msg string, err error)
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
