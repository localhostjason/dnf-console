package casbinx

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/localhostjason/webserver/db"
	"github.com/localhostjason/webserver/util"
)

var E *casbin.Enforcer // 作为全局访问

type CasBin struct {
	File string
}

func NewCasBin() *CasBin {
	cfg, _ := GetCasbinConfig()
	return &CasBin{File: cfg.File}
}

func (c *CasBin) Run() error {
	if !util.PathExists(c.File) {
		createCasbinConfigFile(c.File)
	}

	dbx := db.DB
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(dbx, &CasbinRule{})

	var enforcer *casbin.Enforcer
	enforcer, err = casbin.NewEnforcer(c.File, adapter)

	err = enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	E = enforcer
	return nil
}

func (c *CasBin) Clear(v int, p ...string) bool {
	ok, _ := E.RemoveFilteredPolicy(v, p...)
	return ok
}

func (c *CasBin) Get() [][]string {
	policy := E.GetPolicy()
	return policy
}
