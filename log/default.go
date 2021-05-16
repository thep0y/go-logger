/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: default.go
 * @Created: 2021-05-16 09:39:17
 * @Modified: 2021-05-16 19:54:27
 */

package log

var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
    defaultLogger.depth = 2
}

// Trace 默认 Trace 方法
func Trace(v ...interface{}) {
	defaultLogger.Trace(v...)
}

// Tracef 默认 Tracef 方法
func Tracef(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v...)
}

// Info 默认 Info 方法
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof 默认 Infof 方法
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Debug 默认 Debug 方法
func Debug(v ...interface{}) {
    defaultLogger.Debug(v...)
}

// Debugf 默认 Debugf 方法
func Debugf(format string, v ...interface{}) {
    defaultLogger.Debugf(format, v...)
}

// Warn 默认 Warn 方法
func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

// Warnf 默认 Warnf 方法
func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

// Error 默认 Error 方法
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf 默认 Errorf 方法
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Fatal 默认 Fatal 方法
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Fatalf 默认 Fatalf 方法
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}
