package main

import (
	"dtp/core/conf"
	"dtp/core/http"
)

func main() {
	conf.Start(".")

	sock := "/uds-call-go.sock"
	appConf := conf.Get("app", "app").(conf.AppConf)
	uds := appConf.Path + sock

	go conf.Load(".")

	go http.Go(uds)

	http.Start()
}
