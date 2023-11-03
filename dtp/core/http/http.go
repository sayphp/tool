package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"dtp/core/call"
	"dtp/core/conf"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// * 服务启动
func Start() {
	http.HandleFunc("/", Router)
	http.ListenAndServe(":80", nil)
}

// * 路由
func Router(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.RequestURI())
	routerList := conf.List("router")

	key := strings.Replace(u.Path, "/dtp/", "", 1)

	//fmt.Println("%+v, %+v, %+v", u, key, conf.List("router"))
	if routerList[key] == nil {
		fmt.Printf("预期外路由:%s\n", u.Path)
		http.NotFound(w, r)
		return
	}
	router := routerList[key].(conf.RouterConf)
	res := conf.Res{
		Code: 0,
		Msg:  "ok",
	}
	switch router.Type {
	case "go":
		call.Go(u.Path, nil, r, w, &res)
	case "php":
		res.Data = call.Php(u.Path, nil, r, w)
	case "python":
		res.Data = call.Py(u.Path, nil, r, w)
	case "shell":
		res.Data = call.Sh(u.Path, nil, r, w)
	case "javascript":
		res.Data = call.Js(u.Path, nil, r, w)
	default:
		fmt.Println("预期外调用:%s", router.Type)
		http.NotFound(w, r)
	}
	s, _ := json.Marshal(res)
	io.WriteString(w, string(s))
}
