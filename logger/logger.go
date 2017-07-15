package logger

import (
	"log"
	"os"
	"sync"

	"github.com/sycdtk/gotools/config"
)

const (
	//日志级别
	ldebug = 1 << 0
	linfo  = 1 << 1

	DEBUG = "DEBUG"
	INFO  = "INFO"
)

type Logger struct {
	*log.Logger
	level int //日志级别
	m     *sync.Mutex
}

//设置日志级别
func (l *Logger) SetLevel(lvl string) {
	if lvl == DEBUG {
		l.level = ldebug
	} else if lvl == INFO {
		l.level = linfo
	}

}

//debug输出，包含info输出
func (l *Logger) Debug(v ...interface{}) {
	if ldebug == l.level&ldebug {
		l.m.Lock()
		l.Logger.SetPrefix("D: ")
		l.Logger.Println(v)
		l.Logger.SetPrefix("   ")
		l.m.Unlock()
	}
}

//info输出
func (l *Logger) Info(v ...interface{}) {
	if linfo == l.level&linfo || ldebug == l.level&ldebug {
		l.m.Lock()
		l.Logger.SetPrefix("I: ")
		l.Logger.Println(v)
		l.Logger.SetPrefix("   ")
		l.m.Unlock()
	}
}

var mylogger *Logger
var once sync.Once

func NewLogger() *Logger {

	//仅执行一次，单例
	once.Do(func() {

		filePath := config.Read("log", "path")
		logLevel := config.Read("log", "level")

		println(filePath)

		var logstd *log.Logger

		if len(filePath) != 0 { //日志文件

			file, err := os.Create(filePath)
			if err != nil {
				log.Panicln("创建日志失败!", err)
			}

			logstd = log.New(file, "", log.LstdFlags|log.Lshortfile) //构建默认log对象

		} else { //标准输出
			logstd = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile) //构建默认log对象
		}

		mylogger = &Logger{logstd, ldebug, new(sync.Mutex)} //默认级别

		if len(logLevel) > 0 {
			mylogger.SetLevel(logLevel) //设置日志级别
		}

		mylogger.SetFlags(log.LstdFlags | log.Lshortfile) //设置输出格式
	})

	return mylogger

}
