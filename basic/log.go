/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: log.go
 * @Created: 2021-05-16 09:39:17
 * @Modified: 2021-05-16 10:31:46
 */

package basic

import (
    "io"
    "os"
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

var logger Logger = &EmptyLogger{}

type EmptyLogger struct{}

func (l *EmptyLogger) SetLogLevel(LogLevel) {}

func (l *EmptyLogger) Fatal(v ...interface{}) { os.Exit(1) }

func (l *EmptyLogger) Fatalf(format string, v ...interface{}) { os.Exit(1) }

func (l *EmptyLogger) Error(v ...interface{}) {}

func (l *EmptyLogger) Errorf(format string, v ...interface{}) {}

func (l *EmptyLogger) Warn(v ...interface{}) {}

func (l *EmptyLogger) Warnf(format string, v ...interface{}) {}

func (l *EmptyLogger) Info(v ...interface{}) {}

func (l *EmptyLogger) Infof(format string, v ...interface{}) {}

func (l *EmptyLogger) Debug(v ...interface{}) {}

func (l *EmptyLogger) Debugf(format string, v ...interface{}) {}

func (l *EmptyLogger) Trace(v ...interface{}) {}

func (l *EmptyLogger) Tracef(format string, v ...interface{}) {}

func (l *EmptyLogger) SetOutput(w io.Writer) {}

func Error(v ...interface{}) {
    logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
    logger.Errorf(format, v...)
}

func Warn(v ...interface{}) {
    logger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
    logger.Warnf(format, v...)
}

func Info(v ...interface{}) {
    logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
    logger.Infof(format, v...)
}

func Debug(v ...interface{}) {
    logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
    logger.Debugf(format, v...)
}

func Trace(v ...interface{}) {
    logger.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
    logger.Tracef(format, v...)
}

func Fatal(v ...interface{}) {
    logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
    logger.Fatalf(format, v...)
}

func SetLogLevel(level LogLevel) {
    logger.SetLogLevel(level)
}

func SetOutput(w io.Writer) {
    logger.SetOutput(w)
}

func RegisterLogger(l Logger) {
    logger = l
}
