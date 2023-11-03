package main

import (
	"dtp/core/conf"
	"dtp/core/http"
)

func main() {

	udsConf := conf.Get("app", "uds").(conf.AppConf)

	go conf.Load()

	go http.Go(udsConf.UdsGo)

	http.Start()
}
