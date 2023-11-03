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
				e, _ := err.(reflect.Value)
				res, _ := json.Marshal(e.Interface())
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
			fmt.Println("yaegi eval fail", err, content)
			return
		}

		v, err := i.Eval("plugin.run")

		if err != nil {
			fmt.Println(err)
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
