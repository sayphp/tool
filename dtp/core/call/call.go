package call

import (
	"net/http"
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

}

// * 失败
func Error(w http.ResponseWriter) {

}
