package plugin

import (
	"net/http"

	"dtp/core/call"
	"dtp/core/conf"
)

type Conf struct {
	App    conf.AppConf    `json:"app"`
	Db     conf.DbConf     `json:"db"`
	Notice conf.NoticeConf `json:"notice"`
	Router conf.RouterConf `json:"router"`
	Task   conf.TaskConf   `json:"task"`
}

func run(r *http.Request, w http.ResponseWriter) interface{} {
	var c Conf
	call.Go("/dtp/demo/conf", nil, r, w, &c) //call /dtp/demo/conf 返回结果
	return c
}
