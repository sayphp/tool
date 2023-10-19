package main

import (
	//"dtp/core/http"
	"dtp/core/conf"
)

const (
	PATH = "/say/code/liuxiao/say/dtp"
)

func main() {

	conf.Load(PATH)
	//http.Start()
}
