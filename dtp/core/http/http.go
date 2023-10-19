package http

import (
	"dtp/core/call"
	"dtp/core/conf"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//* 服务启动
func Start() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}

//* 路由
func router(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.RequestURI())
	appConf := conf.Get("app", "app").(conf.AppConf)
	routerList := conf.List("router")

	key := strings.Replace(u.Path, "/dtp/", "", 1)

	//fmt.Println("%+v, %+v, %+v", u, key, conf.List("router"))
	if routerList[key] == nil {
		fmt.Println("预期外路由:%s", u.Path)
		http.NotFound(w, r)
		return
	}
	router := routerList[key].(conf.RouterConf)
	file := appConf.Path + "/plugin/" + router.Type + router.Path
	switch router.Type {
	case "go":
		call.Go(file, r, w)
	case "php":
		call.Php(file, r, w)
	case "python":
		call.Py(file, r, w)
	case "shell":
		call.Sh(file, r, w)
	case "javascript":
		call.Js(file, r, w)
	default:
		fmt.Println("预期外调用:%s", router.Type)
		http.NotFound(w, r)
	}
}

//* 成功
func Succ(w http.ResponseWriter) {

}

//* 失败
func Fail(w http.ResponseWriter) {

}
