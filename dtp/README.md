# 开发工具平台 development tool platform

## 设计思路

* 插件管理
* 跨语言调用


## 目录结构

```bash
|- conf
|- core
|- plugin
main.go 启动文件
```

## 版本依赖

```shell
go version go1.21.3 linux/amd64
```

## 框架使用

```go

//框架内部获取配置
// conf.Get("配置类型", "配置文件").(强制类型转换)
//* 配置类型
mode = {
    "db",//数据库，连接配置
    "notice",//通知，钉钉、邮件等服务使用
    "router",//路由，http服务使用
    "task",//任务，定时任务使用
}
//配置类型
conf = {
    DbConf,//存储配置
    NoticeConf,//通知配置
    RouterConf,//路由配置
    TaskConf, //任务配置
}
acc9 := conf.Get("db", "acc/9").(conf.DbConf) //请注意强转类型

```