package plugin

import (
	"fmt"
	"net/http"

	"dtp/core/conf"
)

type Test struct {
	Key string       `json:"key"`
	Val int          `json:"val"`
	App conf.AppConf `json:"app"`
}

func run(r *http.Request, w http.ResponseWriter) interface{} {
	fmt.Println("call /dtp/test/go:ABC")
	conf.Start(".")
	fmt.Println("%+v", conf.Get("app", "app"))
	return Test{
		Key: "ABC",
		Val: 456,
		App: conf.Get("app", "app").(conf.AppConf),
	}
}
