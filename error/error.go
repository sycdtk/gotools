package err

import (
	"log"
)

//检查错误
func CheckError(err error) {
	if err != nil {
		log.Panicln("异常：", err)
	}
}
