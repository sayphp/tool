package util

import (
	"dtp/core/conf"
	"fmt"
)

type ErrorInfo struct {
	Line string `json:"line"`
	Code string `json:"code"`
}

// * 失败
func Error(code int, opt ...interface{}) {
	var msg string
	if ErrorMap[code] != "" {
		msg = ErrorMap[code]
	}
	var data interface{}
	fmt.Printf("%+v\n", opt)
	l := len(opt)
	if l > 0 {
		if opt[0] != nil {
			if msg != "" {
				msg = msg + " -"
			}
			msg = msg + opt[0].(string)
		}

		if l == 2 {
			data = opt[1]
		}
	}

	if msg == "" {
		msg = "unknow error"
	}
	e := conf.Res{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	panic(e)
}

const (
	E_CORE   = 1
	E_PARAM  = 2
	E_CALL   = 3
	E_RPC    = 4
	E_INTER  = 5
	E_SERVER = 6
	E_MYSQL  = 201
	E_UNKNOW = 999
)

var ErrorMap = map[int]string{
	E_CORE:   "core error",
	E_PARAM:  "param error",
	E_CALL:   "call error",
	E_RPC:    "rpc error",
	E_MYSQL:  "mysql error",
	E_INTER:  "interpreter error",
	E_SERVER: "server error",
	E_UNKNOW: "unknow error",
}
