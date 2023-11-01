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

var conf map[string]map[string]interface{}

// * 获取
func Get(mode string, name string) interface{} {
	return conf[mode][name]
}

// * 列表
func List(mode string) map[string]interface{} {
	return conf[mode]
}

// 初始化结构
func build(path string) {
	conf = make(map[string]map[string]interface{})
	_ = filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}
		mode, name := getKey(p, "conf")
		if conf[mode] == nil {
			conf[mode] = make(map[string]interface{})
		}

		content, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Println("配置文件读取异常:%s", err)
			return nil
		}
		//fmt.Println("%+v", content)

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

// * 初始化
func Start(path string) {
	build(path + "/conf")
}

// * 动态加载
func Load(path string) {
	for {
		build(path + "/conf")
		time.Sleep(10 * time.Second)
	}

}

func getKey(path string, root string) (mode string, name string) {
	tmp := strings.Replace(path, root, "", 1)
	regp, err := regexp.Compile("/(.*?)/")
	if err != nil {
		return
	}
	mode = regp.FindString(tmp)
	name = strings.Replace(tmp, mode, "", 1)
	return strings.Replace(mode, "/", "", 2), strings.Replace(name, ".json", "", 1)
}
