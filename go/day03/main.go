package main

import (
	"logger/logger"
	"time"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	m["log_split_type"] = "size"
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}

	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("user server is running, :/Users/hxadmin/workspace/study-notes/go/logger")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "/Users/hxadmin/workspace/study-notes/go/logger/", "user_server", "debug")
	Run()
}
