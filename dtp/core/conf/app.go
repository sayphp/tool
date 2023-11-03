package conf

type AppConf struct {
	App    string `json:"app"`
	Ver    string `json:"ver"`
	Auth   string `json:"auth"`
	Path   string `json:"path"`
	Gopath string `json:"gopath"`
	UdsGo  string `json:"udsGo"`
	UdsPhp string `json:"udsPhp"`
}
