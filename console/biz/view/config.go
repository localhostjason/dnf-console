package view

import "github.com/localhostjason/webserver/server/config"

const _key = "api"

type ConfigView struct {
	CORS bool `json:"cors"`
}

func init() {
	c := ConfigView{CORS: true}
	_ = config.RegConfig(_key, c)
}

func GetConfig() (ConfigView, error) {
	var c ConfigView
	err := config.GetConfig(_key, &c)
	return c, err
}
