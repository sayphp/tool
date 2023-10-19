package main

import (
	"dtp/core/conf"
	"dtp/core/http"
)

func main() {
	go conf.Load(".")
	http.Start()
}
