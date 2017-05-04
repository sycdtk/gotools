package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	//是否写日志文件
	writeFile = false
	//日志文件名称
	logFile = "logger.log"
	//日志级别
	L_DEBUG = 1 << 0
	L_INFO  = 1 << 1
	L_ERROR = 1 << 2
)

//默认级别，显示info和error
var level = L_INFO | L_ERROR

func init() {
	if writeFile {
		f, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		os.Stdout = f
		os.Stderr = f
	}
}

//设置日志级别，例如：logger.L_INFO | logger.L_ERROR
func SetLevel(l int) {
	level = l
}

func Debug(msg string) {
	if L_DEBUG == level&L_DEBUG {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[DEBUG]:", msg)
	}
}

func Info(msg string) {
	if L_INFO == level&L_INFO {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[INFO ]:", msg)
	}
}

func Error(msg string) {
	if L_ERROR == level&L_ERROR {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[ERROR]:", msg)
	}
}
