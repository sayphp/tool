package plugin

import (
	"net/http"

	"dtp/core/call"
)

func run(r *http.Request, w http.ResponseWriter) interface{} {
	call.Error(999, "未知错误")

	return c
}
