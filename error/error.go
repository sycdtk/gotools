package error

import (
	"errors"

	"github.com/sycdtk/gotools/logger"
)

//创建异常
func NewErr(errStr string) (obj interface{}, err error) {
	logger.Info(errStr)
	return nil, errors.New(errStr)
}

//检查错误
func CheckErr(err error) {
	if err != nil {
		log.Panicln("异常!", err)
	}
}
