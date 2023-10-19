package plugin

import (
	"fmt"
	"net/http"
)

func run(r *http.Request, w http.ResponseWriter) string {
	fmt.Println("ABC")

	return "ABC"
}
