package plugin

import (
	"dtp/core/call"
	"fmt"
	"net/http"
)

func run(r *http.Request, w http.ResponseWriter) string {
	fmt.Println("This is Go Run Test!")
	ret := call.Go("/api/abc.go", r, w)
	return "This is a go::call test :" + ret + " -- succ!"
}
