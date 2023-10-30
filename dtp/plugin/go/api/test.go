package plugin

import (
	"fmt"
	"net/http"

	"dtp/core/call"
)

func run(r *http.Request, w http.ResponseWriter) interface{} {
	fmt.Println("This is Go Run Test!")
	ret := call.Go("/dtp/test/test", nil, r, w)
	return ret
}
