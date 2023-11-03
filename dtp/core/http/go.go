package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"dtp/core/conf"
	"dtp/core/util"

	"yaegi/interp"
	"yaegi/stdlib"
)

func Go(uds string) {
	fmt.Println(uds, "listen start")
	os.Remove(uds)

	l, e := net.Listen("unix", uds)
	if e != nil {
		fmt.Println(uds, "listen error ", e)
		os.Exit(1)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				//fmt.Printf("err:%+v\ntype:%s\n", err, reflect.TypeOf(err).Name())
				var res []byte
				if reflect.TypeOf(err).Name() == "Res" {
					res, _ = json.Marshal(err)
				} else {
					e, _ := err.(reflect.Value)
					res, _ = json.Marshal(e.Interface())
				}
				io.WriteString(w, string(res))
			}
		}()
		u, _ := url.Parse(r.URL.RequestURI())
		appConf := conf.Get("app", "app").(conf.AppConf)
		routerList := conf.List("router")

		key := strings.Replace(u.Path, "/dtp/", "", 1)

		if routerList[key] == nil {
			fmt.Println("预期外路由:%s", u.Path)
			http.NotFound(w, r)
			return
		}
		router := routerList[key].(conf.RouterConf)
		file := appConf.Path + "/plugin/" + router.Type + router.Path
		src, _ := os.ReadFile(file)
		content := string(src)
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		_, err := i.Eval(content)
		if err != nil {
			d := util.ErrorInfo{
				Line: err.Error(),
				Code: content,
			}
			util.Error(5, "yaegi Eval() failed", d)
			return
		}

		v, err := i.Eval("plugin.run")

		if err != nil {
			util.Error(5, "plugin.run() failed", err)
		}

		plugin := v.Interface().(func(r *http.Request, w http.ResponseWriter) interface{})
		o := plugin(r, w)
		ret := conf.Res{
			Code: 0,
			Msg:  "ok",
			Data: o,
		}
		s, _ := json.Marshal(ret)
		io.WriteString(w, string(s))

	})

	server := http.Server{
		Handler: m,
	}

	server.Serve(l)
}
