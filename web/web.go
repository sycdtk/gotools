package web

import (
	"io"
	"net/http"
	"sync"
)

var once sync.Once
var route map[string]func(http.ResponseWriter, *http.Request) //路由集合
var staticRoute map[string][]string                           //静态文件路由集合

//构建默认handler
type webHandler struct {
}

//默认handler，路由选择
func (*webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "/test" == r.URL.String() {
		test(w, r)
		return
	}
	io.WriteString(w, "default") //默认页面
}

type WebServer struct {
	*http.Server
}

//构建函数
func NewWebServer() *WebServer {

	//TODO 读取配置文件

	ws := &WebServer{&http.Server{
		Addr:    ":8080",
		Handler: &webHandler{},
	}}

	return ws
}

//路由注册函数
func RouterRegister(path string) {
	once.Do(func() {

		route = make(map[string]func(http.ResponseWriter, *http.Request)) //路由注册集合构建

		//静态路由注册函数
	})

	route[path] = f
}
