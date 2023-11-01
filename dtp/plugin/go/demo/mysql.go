package plugin

import (
	"net/http"

	"dtp/core/conf"
)

func run(r *http.Request, w http.ResponseWriter) interface{} {
	conf.Start(".")                                //初始化配置 使用前必须
	dbConf := conf.Get("db", "demo").(conf.DbConf) //mysql配置

	return "TODO"
}
