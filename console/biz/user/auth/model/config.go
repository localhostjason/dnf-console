package model

import "github.com/localhostjason/webserver/server/config"

const _key = "auth"

func regConfig() {
	c := ConfigAuth{
		Realm:      "test zone",
		Secret:     "f450a7bdbde3416d22474b9fdc2a3636",
		IDKey:      "username",
		Timeout:    12 * 3600,
		MaxRefresh: 3600,
	}
	_ = config.RegConfig(_key, c)
}

type ConfigAuth struct {
	Realm      string `json:"realm"`
	Secret     string `json:"secret"`
	IDKey      string `json:"id_key"`
	Timeout    int    `json:"timeout"`
	MaxRefresh int    `json:"max_refresh"`
}

func GetConfig() (ConfigAuth, error) {
	var c ConfigAuth
	err := config.GetConfig(_key, &c)
	return c, err
}

func init() {
	regConfig()
}
