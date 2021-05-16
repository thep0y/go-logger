/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: color.go
 * @Created: 2021-05-16 09:05:06
 * @Modified: 2021-05-16 10:39:38
 */

package color

import (
	"runtime"

	"github.com/thep0y/go-logger/basic/buffer"
)

// ColorBuffer 颜色缓冲用来保存颜色
type ColorBuffer struct {
	buffer.Buffer
}

var (
	colorOff    = []byte("\033[0m")
	colorRed    = []byte("\033[0;31m")
	colorGreen  = []byte("\033[0;32m")
	colorOrange = []byte("\033[0;33m")
	colorBlue   = []byte("\033[0;34m")
	colorPurple = []byte("\033[0;35m")
	colorCyan   = []byte("\033[0;36m")
	colorGray   = []byte("\033[0;37m")
)

func init() {
	if runtime.GOOS == "windows" {
		colorOff = []byte("")
		colorRed = []byte("")
		colorGreen = []byte("")
		colorOrange = []byte("")
		colorBlue = []byte("")
		colorPurple = []byte("")
		colorCyan = []byte("")
		colorGray = []byte("")
	}
}

// Off 不用颜色
func (cb *ColorBuffer) Off() {
	cb.Append(colorOff)
}

// Red 用红色
func (cb *ColorBuffer) Red() {
	cb.Append(colorRed)
}

// Green 用绿色
func (cb *ColorBuffer) Green() {
	cb.Append(colorGreen)
}

// Orange 用橙色
func (cb *ColorBuffer) Orange() {
	cb.Append(colorOrange)
}

// Blue 用蓝色
func (cb *ColorBuffer) Blue() {
	cb.Append(colorBlue)
}

// Purple 用紫色
func (cb *ColorBuffer) Purple() {
	cb.Append(colorPurple)
}

// Cyan 用青色
func (cb *ColorBuffer) Cyan() {
	cb.Append(colorCyan)
}

// Gray 用灰色
func (cb *ColorBuffer) Gray() {
	cb.Append(colorGray)
}

// 将颜色和 off 字节与实际数据混合在一起
func mixer(data []byte, color []byte) []byte {
	var result []byte
	return append(append(append(result, color...), data...), colorOff...)
}

// Red 将 data 染成红色
func Red(data []byte) []byte {
	return mixer(data, colorRed)
}

// Green 将 data 染成绿色
func Green(data []byte) []byte {
	return mixer(data, colorGreen)
}

// Orange 将 data 染成橙色
func Orange(data []byte) []byte {
	return mixer(data, colorOrange)
}

// Blue 将 data 染成蓝色
func Blue(data []byte) []byte {
	return mixer(data, colorBlue)
}

// Purple 将 data 染成紫色
func Purple(data []byte) []byte {
	return mixer(data, colorPurple)
}

// Cyan 将 data 染成青色
func Cyan(data []byte) []byte {
	return mixer(data, colorCyan)
}

// Gray 将 data 染成灰色
func Gray(data []byte) []byte {
	return mixer(data, colorGray)
}
