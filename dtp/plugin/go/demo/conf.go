package plugin

import (
	"fmt"
	"net/http"

	"dtp/core/conf"
)

type Conf struct {
	App    conf.AppConf      `json:"app"`
	Db     conf.DbConf       `json:"db"`
	Notice conf.NoticeConf   `json:"notice"`
	Router conf.RouterConf   `json:"router"`
	Task   conf.TaskConf     `json:"task"`
	List   map[string]string `json:"list"`
}

func run(r *http.Request, w http.ResponseWriter) interface{} {

	//获取配置
	appConf := conf.Get("app", "app").(conf.AppConf)                //应用配置
	dbConf := conf.Get("db", "demo").(conf.DbConf)                  //mysql配置
	noticeConf := conf.Get("notice", "demo").(conf.NoticeConf)      //通知配置
	routerConf := conf.Get("router", "demo/conf").(conf.RouterConf) //路由配置
	taskConf := conf.Get("task", "demo").(conf.TaskConf)            //任务配置
	routerList := conf.List("router")                               //获取配置列表
	c := Conf{
		App:    appConf,
		Db:     dbConf,
		Notice: noticeConf,
		Router: routerConf,
		Task:   taskConf,
	}
	//配置列表使用方法
	c.List = make(map[string]string)
	for k, v := range routerList {
		fmt.Printf("%T:%+v\n", v, v)
		c.List[k] = v.(conf.RouterConf).Path
	}

	return c
}
