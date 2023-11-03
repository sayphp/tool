package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const PATH = "conf"

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Caller struct {
	File string `json:"file"`
	Line int    `json:"line"`
	Name string `json:"name"`
}

var conf map[string]map[string]interface{}

// * 获取
func Get(mode string, name string) interface{} {
	if conf[mode][name] == nil {
		build(mode, name)
	}
	return conf[mode][name]
}

// * 列表
func List(mode string) map[string]interface{} {
	if conf[mode] == nil {
		build(mode)
	}
	return conf[mode]
}

// 初始化结构
func build(opt ...string) {
	path := PATH
	l := len(opt)
	if l > 1 { // mode
		path = path + "/" + opt[0]
		if l == 2 { // mode
			path = path + "/" + opt[1] + ".json"
		}
	}

	conf = make(map[string]map[string]interface{})
	_ = filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}
		mode, name := getKey(p)
		if conf[mode] == nil {
			conf[mode] = make(map[string]interface{})
		}

		content, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Println("配置文件读取异常:%s", err)
			return nil
		}
		//fmt.Printf("content:%+s\n", content)

		if string(content) == "" {
			fmt.Println("空配置文件:%s", p)
			return nil
		}
		switch mode {
		case "db":
			var c DbConf
			_ = json.Unmarshal(content, &c)
			conf[mode][name] = c
		case "notice":
			var c NoticeConf
			_ = json.Unmarshal(content, &c)
			conf[mode][name] = c
		case "router":
			var c RouterConf
			_ = json.Unmarshal(content, &c)
			conf[mode][name] = c
		case "task":
			var c TaskConf
			_ = json.Unmarshal(content, &c)
			conf[mode][name] = c
		case "app":
			var c AppConf
			_ = json.Unmarshal(content, &c)
			conf[mode][name] = c
			//fmt.Printf("AppConf:%+v\n", c)
		default:
			fmt.Println("预期外配置类型:%s", mode)
			return nil
		}
		return nil
	})
	//fmt.Printf("conf:%+v\n", conf)
}

// * 初始化 - 全部配置
func Init() {
	build()
}

// * 动态加载
func Load() {
	for {
		build()
		time.Sleep(10 * time.Second)
	}

}

func getKey(path string) (mode string, name string) {
	tmp := strings.Replace(path, "./", "", 1) //解决指定路径问题
	rep := strings.Replace(PATH, "./", "", 1)
	tmp = strings.Replace(tmp, rep, "", 1)
	//fmt.Printf("path:%+v\n", path)
	regp, err := regexp.Compile("/(.*?)/")
	if err != nil {
		return
	}
	mode = regp.FindString(tmp)
	name = strings.Replace(tmp, mode, "", 1)

	mode = strings.Replace(mode, "/", "", 2)
	name = strings.Replace(name, ".json", "", 1)
	return mode, name
}
