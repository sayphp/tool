package plugin

import (
	"net/http"

	"dtp/core/util"
)

func run(r *http.Request, w http.ResponseWriter) interface{} {
	util.Error(1)
	util.Error(99, "error test succ")

	return "error test failed"
}
