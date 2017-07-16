package web

import (
	"github.com/sycdtk/gotools/web/session"
	"github.com/sycdtk/gotools/web/session/memory"
)

var SessionManager *session.SessionManager

//然后在init函数中初始化
func init() {
	session.Register("memory", memory.NewMemProvider())
	SessionManager, _ = session.NewSessionManager("memory", "gosessionID", 30, 100)

	go SessionManager.GC()
}
