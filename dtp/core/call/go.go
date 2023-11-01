package call

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"dtp/core/conf"
)

func Go(path string, args interface{}, r *http.Request, w http.ResponseWriter, ret interface{}) {
	conf.Start(".")
	sock := "/uds-call-go.sock"
	appConf := conf.Get("app", "app").(conf.AppConf)
	uds := appConf.Path + sock
	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", uds)
			},
		},
	}

	var resp *http.Response
	resp, _ = httpc.Get("http://localhost" + path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("body:%+v\n", string(body))

	json.Unmarshal(body, &ret)
}
