package plugin

import (
	"net/http"

	"dtp/core/conf"
)

func run(r *http.Request, w http.ResponseWriter) interface{} {

	dbConf := conf.Get("db", "demo").(conf.DbConf) //mysql配置

	return "TODO"
}
