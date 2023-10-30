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

// * 成功
func Succ(w http.ResponseWriter) {

}

// * 失败
func Fail(w http.ResponseWriter) {

}
