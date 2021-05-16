/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: log.go
 * @Created: 2021-05-16 09:39:17
 * @Modified: 2021-05-16 10:51:53
 */

package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	log "github.com/thep0y/go-logger/basic"
	"github.com/thep0y/go-logger/basic/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Logger 日志的底层存储
type Logger struct {
	mu        sync.RWMutex
	color     bool
	out       io.Writer
	debug     bool
	timestamp bool
	quiet     bool
	buf       color.ColorBuffer
	logLevel  log.LogLevel
}

// Prefix 日志信息的前缀和颜色
type Prefix struct {
	Plain []byte
	Color []byte
	File  bool
}

type FdWriter interface {
	io.Writer
	Fd() uintptr
}

// NewLogger 返回一个输出到指定位置的自动染色的日志实例
func NewLogger() *Logger {
	logger := &Logger{
		color:     terminal.IsTerminal(int(os.Stdout.Fd())),
		out:       os.Stdout,
		timestamp: true,
	}
	log.RegisterLogger(logger)
	return logger
}

var (
	// 日志前缀
	plainFatal = []byte("[FATAL] ")
	plainError = []byte("[ERROR] ")
	plainWarn  = []byte("[WARN]  ")
	plainInfo  = []byte("[INFO]  ")
	plainDebug = []byte("[DEBUG] ")
	plainTrace = []byte("[TRACE] ")

	// Fatal 前缀
	FatalPrefix = Prefix{
		Plain: plainFatal,
		Color: color.Red(plainFatal),
		File:  true,
	}

	// Error 前缀
	ErrorPrefix = Prefix{
		Plain: plainError,
		Color: color.Red(plainError),
		File:  true,
	}

	// Warn 前缀
	WarnPrefix = Prefix{
		Plain: plainWarn,
		Color: color.Orange(plainWarn),
	}

	// Info 前缀
	InfoPrefix = Prefix{
		Plain: plainInfo,
		Color: color.Green(plainInfo),
	}

	// Debug 前缀
	DebugPrefix = Prefix{
		Plain: plainDebug,
		Color: color.Purple(plainDebug),
		File:  true,
	}

	// Trace 前缀
	TracePrefix = Prefix{
		Plain: plainTrace,
		Color: color.Cyan(plainTrace),
	}
)

// SetLogLevel 设置日志等级
func (l *Logger) SetLogLevel(level log.LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

// SetOutput 设置日志输出位置
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = false
	if fdw, ok := w.(FdWriter); ok {
		l.color = terminal.IsTerminal(int(fdw.Fd()))
	}
	l.out = w
}

// WithColor 启用日志染色
func (l *Logger) WithColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = true
	return l
}

// WithoutColor 关闭日志染色
func (l *Logger) WithoutColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = false
	return l
}

// WithDebug 打开调试模式，用来调试和追踪
func (l *Logger) WithDebug() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debug = true
	return l
}

// WithoutDebug 关闭调试模式
func (l *Logger) WithoutDebug() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debug = false
	return l
}

// IsDebug 检查调式模式的状态
func (l *Logger) IsDebug() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.debug
}

// WithTimestamp 开启时间戳输出
func (l *Logger) WithTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = true
	return l
}

// WithoutTimestamp 关闭时间戳输出
func (l *Logger) WithoutTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = false
	return l
}

// Quiet 关闭所有日志输出
func (l *Logger) Quiet() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.quiet = true
	return l
}

// NoQuiet 开启所有日志输出
func (l *Logger) NoQuiet() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.quiet = false
	return l
}

// IsQuiet 检查当前日志输出状态
func (l *Logger) IsQuiet() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.quiet
}

// Output 输出实际数据
func (l *Logger) Output(depth int, prefix Prefix, data string) error {
	// 如果当前禁止输出日志，返回 nil，不输出任何信息
	if l.IsQuiet() {
		return nil
	}

	now := time.Now()

	// 追踪文件和代码行号的临时存储
	var file string
	var line int
	var fn string

	// 检查指定的前缀是否需要包含文件日志记录
	if prefix.File {
		var ok bool
		var pc uintptr

		// 获取调用者的文件名和代码行号
		if pc, file, line, ok = runtime.Caller(depth + 1); !ok {
			file = "<unknown file>"
			fn = "<unknown function>"
			line = 0
		} else {
			file = filepath.Base(file)
			fn = runtime.FuncForPC(pc).Name()
		}
	}
	// 共享缓冲区上锁
	l.mu.Lock()
	defer l.mu.Unlock()
	// 重置缓冲区
	l.buf.Reset()
	// 向缓冲区写前缀
	if l.color {
		l.buf.Append(prefix.Color)
	} else {
		l.buf.Append(prefix.Plain)
	}
	// 检查是否显示时间戳
	if l.timestamp {
		// 如果开启颜色，将时间戳染色
		if l.color {
			l.buf.Blue()
		}
		// 输出日期的时间
		year, month, day := now.Date()
		l.buf.AppendInt(year, 4)
		l.buf.AppendByte('/')
		l.buf.AppendInt(int(month), 2)
		l.buf.AppendByte('/')
		l.buf.AppendInt(day, 2)
		l.buf.AppendByte(' ')
		hour, min, sec := now.Clock()
		l.buf.AppendInt(hour, 2)
		l.buf.AppendByte(':')
		l.buf.AppendInt(min, 2)
		l.buf.AppendByte(':')
		l.buf.AppendInt(sec, 2)
		l.buf.AppendByte(' ')
		// 如果开启颜色，则关闭颜色
		if l.color {
			l.buf.Off()
		}
	}
	// 如果开启 file ，添加调用者的文件名和行号
	if prefix.File {
		// 染色
		if l.color {
			l.buf.Orange()
		}
		// 打印文件名和行号
		l.buf.Append([]byte(fn))
		l.buf.AppendByte(':')
		l.buf.Append([]byte(file))
		l.buf.AppendByte(':')
		l.buf.AppendInt(line, 0)
		l.buf.AppendByte(' ')
		// 关闭颜色
		if l.color {
			l.buf.Off()
		}
	}
	// 打印实际数据
	l.buf.Append([]byte(data))
	if len(data) == 0 || data[len(data)-1] != '\n' {
		l.buf.AppendByte('\n')
	}
	// 冲洗缓冲区内的数据到输出位置
	_, err := l.out.Write(l.buf.Buffer)
	return err
}

// Fatal 打印 fatal 日志，并以状态码 1 退出当前程序
func (l *Logger) Fatal(v ...interface{}) {
	if l.logLevel <= log.FatalLevel {
		l.Output(1, FatalPrefix, fmt.Sprintln(v...))
	}
	os.Exit(1)
}

// Fatalf 根据指定的格式打印 fatal 日志，并以状态码 1 退出当前程序
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.logLevel <= log.FatalLevel {
		l.Output(1, FatalPrefix, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

// Error 打印 error 日志
func (l *Logger) Error(v ...interface{}) {
	if l.logLevel <= log.ErrorLevel {
		l.Output(1, ErrorPrefix, fmt.Sprintln(v...))
	}
}

// Errorf 根据指定格式打印 error 日志
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.logLevel <= log.ErrorLevel {
		l.Output(1, ErrorPrefix, fmt.Sprintf(format, v...))
	}
}

// Warn 打印 warning 日志
func (l *Logger) Warn(v ...interface{}) {
	if l.logLevel <= log.WarnLevel {
		l.Output(1, WarnPrefix, fmt.Sprintln(v...))
	}
}

// Warnf 根据指定格式打印 warning 日志
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.logLevel <= log.WarnLevel {
		l.Output(1, WarnPrefix, fmt.Sprintf(format, v...))
	}
}

// Info 打印 info 日志
func (l *Logger) Info(v ...interface{}) {
	if l.logLevel <= log.InfoLevel {
		l.Output(1, InfoPrefix, fmt.Sprintln(v...))
	}
}

// Infof 根据指定格式打印 info 日志
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.logLevel <= log.InfoLevel {
		l.Output(1, InfoPrefix, fmt.Sprintf(format, v...))
	}
}

// Debug 打印 debug 日志
func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel == log.AllLevel {
		l.Output(1, DebugPrefix, fmt.Sprintln(v...))
	}
}

// Debugf 根据指定格式打印 debug 日志
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.logLevel == log.AllLevel {
		l.Output(1, DebugPrefix, fmt.Sprintf(format, v...))
	}
}

// Trace 如果开启调式模式则打印 trace 日志
func (l *Logger) Trace(v ...interface{}) {
	if l.logLevel == log.AllLevel {
		l.Output(1, TracePrefix, fmt.Sprintln(v...))
	}
}

// Tracef 如果开启调式模式则根据指定格式打印 trace 日志
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.logLevel == log.AllLevel {
		l.Output(1, TracePrefix, fmt.Sprintf(format, v...))
	}
}
