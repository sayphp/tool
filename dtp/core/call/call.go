package call

import (
	"net/http"

	"dtp/core/conf"
)

func Php(path string, args interface{}, r *http.Request, w http.ResponseWriter) string {
	return "call::php:" + path
}

func Py(path string, args interface{}, r *http.Request, w http.ResponseWriter) string {
	return "call::python:" + path
}

func Sh(path string, args interface{}, r *http.Request, w http.ResponseWriter) string {
	return "call::shell:" + path
}

func Js(path string, args interface{}, r *http.Request, w http.ResponseWriter) string {
	return "call::javascript:" + path
}

// * 转换参数为url参数
func ConvertParams(args interface{}) string {
	return ""
}

// * 失败
func Error(code int, msg string, data interface{}) {
	//_, file, line, _ := runtime.Caller(1)
	//pc, _, _, _ := runtime.Caller(-1)
	//name := runtime.FuncForPC(pc).Name()
	e := conf.Res{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	panic(e)
}
