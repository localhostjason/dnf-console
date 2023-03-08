package casbinx

import (
	"github.com/localhostjason/webserver/server/config"
	"os"
	"path/filepath"
)

const _casbinKey = "casbin"
const casbinText = `[request_definition]
r = sub,obj,act

[policy_definition]
p = sub,obj,act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act
`

func regCasbinConfig() {
	extPath, _ := os.Getwd()
	file := filepath.Join(extPath, "config", "rbac_model.conf")

	c := CasbinConfig{
		File: file,
	}
	_ = config.RegConfig(_casbinKey, c)
}

type CasbinConfig struct {
	File string `json:"file"`
}

func createCasbinConfigFile(file string) {
	_ = os.WriteFile(file, []byte(casbinText), os.ModePerm)
}

func GetCasbinConfig() (CasbinConfig, error) {
	var c CasbinConfig
	err := config.GetConfig(_casbinKey, &c)
	return c, err
}

func init() {
	regCasbinConfig()
}
