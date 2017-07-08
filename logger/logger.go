package logger

import (
	"log"
	"os"
	"sync"

	"github.com/sycdtk/gotools/config"
)

const (
	//日志级别
	Ldebug = 1 << 0
	Linfo  = 1 << 1
)

type Logger struct {
	*log.Logger
	level int //日志级别
	m     *sync.Mutex
}

//设置日志级别
func (l *Logger) SetLevel(lvl int) {
	l.level = lvl
}

//debug输出，包含info输出
func (l *Logger) Debug(v ...interface{}) {
	if Ldebug == l.level&Ldebug {
		l.m.Lock()
		l.Logger.SetPrefix("D: ")
		l.Logger.Println(v)
		l.Logger.SetPrefix("   ")
		l.m.Unlock()
	}
}

//info输出
func (l *Logger) Info(v ...interface{}) {
	if Linfo == l.level&Linfo || Ldebug == l.level&Ldebug {
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

		filePath := config.Read("logger", "path")

		if len(filePath) == 0 {
			filePath = "logger.log" //默认日志文件
		}

		file, err := os.Create(filePath)
		if err != nil {
			log.Panicln("创建日志失败!", err)
		}
		logstd := log.New(file, "", log.LstdFlags|log.Lshortfile)

		mylogger = &Logger{logstd, Ldebug, new(sync.Mutex)} //默认级别
		mylogger.SetOutput(file)                            //设置标准输出
		mylogger.SetFlags(log.LstdFlags | log.Lshortfile)   //设置输出格式
	})

	return mylogger

}
