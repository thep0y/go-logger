/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: log.go
 * @Created: 2021-05-16 09:39:17
 * @Modified: 2021-05-16 19:31:49
 */

package basic

import (
	"io"
)

// LogLevel 日志等级
type LogLevel int

const (
	AllLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	OffLevel
)

// Logger 日志接口
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	SetLogLevel(level LogLevel)
	SetOutput(io.Writer)
}
