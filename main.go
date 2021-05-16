package main

import (
	"strings"

	"github.com/thep0y/go-logger/log"
)

func main() {
	log.Info("这是默认 info 消息")
	log.Infof("这是默认格式化的消息：%s", "info")
	log.Warn("这是默认 warning 消息")
	log.Warnf("这是默认格式化的消息：%s", "warning")
	log.Error("这是默认error 消息")
	log.Errorf("这是默认格式化的消息：%s", "error")
	// log.Fatal("这是默认 fatal 消息")

	println(strings.Repeat("-", 60))

	logger := log.NewLogger()
	logger.Info("这是 info 消息")
	logger.Infof("这是格式化的消息：%s", "info")
	logger.Warn("这是 warning 消息")
	logger.Warnf("这是格式化的消息：%s", "warning")
	logger.Error("这是 error 消息")
	logger.Errorf("这是格式化的消息：%s", "error")
	logger.Fatal("这是 fatal 消息")
}
