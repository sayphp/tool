package main

import (
	"dtp/core/conf"
	"encoding/json"
	"fmt"
)

type Conf struct {
	App        conf.AppConf      `json:"app"`
	Db         conf.DbConf       `json:"db"`
	Notice     conf.NoticeConf   `json:"notice"`
	Router     conf.RouterConf   `json:"router"`
	Task       conf.TaskConf     `json:"task"`
	RouterList interface{}       `json:"routerList"`
	Test       map[string]string `json:"test"`
}

func run() interface{} {

	//获取配置
	appConf := conf.Get("app", "app").(conf.AppConf)                //应用配置
	dbConf := conf.Get("db", "demo").(conf.DbConf)                  //mysql配置
	noticeConf := conf.Get("notice", "demo").(conf.NoticeConf)      //通知配置
	routerConf := conf.Get("router", "demo/conf").(conf.RouterConf) //路由配置
	taskConf := conf.Get("task", "demo").(conf.TaskConf)            //任务配置
	routerList := conf.List("router")                               //获取配置列表
	c := Conf{
		App:        appConf,
		Db:         dbConf,
		Notice:     noticeConf,
		Router:     routerConf,
		Task:       taskConf,
		RouterList: routerList,
	}
	c.Test = make(map[string]string)
	for k, v := range routerList {
		fmt.Printf("%T:%+v\n", v, v)

		c.Test[k] = v.(conf.RouterConf).Path
	}

	fmt.Printf("%T:%+v\n", routerList, routerList)

	return c
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v", err)
			s, _ := json.Marshal(err)
			fmt.Println(string(s))
		}
	}()
	run()
}
