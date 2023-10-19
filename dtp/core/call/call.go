package call

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	//yaegi:tags safe
	"yaegi/interp"
	"yaegi/stdlib"
)

func Go(path string, r *http.Request, w http.ResponseWriter) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("%+v", err)
	}
	src := string(data)
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err = i.Eval(src)
	if err != nil {
		fmt.Println("%+v", err)
	}

	v, err := i.Eval("plugin.run")
	if err != nil {
		fmt.Println("%+v", err)
	}

	Run := v.Interface().(func(r *http.Request, w http.ResponseWriter) string)

	ret := Run(r, w)
	io.WriteString(w, ret)
	return "go dynamic test"
}

func Php(path string, r *http.Request, w http.ResponseWriter) {

}

func Py(path string, r *http.Request, w http.ResponseWriter) {

}

func Sh(path string, r *http.Request, w http.ResponseWriter) {

}

func Js(path string, r *http.Request, w http.ResponseWriter) {

}
