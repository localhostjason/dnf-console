package game_db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DBPools = new(ConnectPool)

func Connect() error {
	c := getDbConfig()
	if !c.Enable {
		fmt.Println("no use db")
		return nil
	}

	if len(c.Mysql) != 0 {
		if err := ConnectWithMysqlConfig(c.Mysql); err != nil {
			return err
		}
	}

	if len(c.Sqlite) != 0 {
		if err := ConnectWithSqliteConfig(c.Sqlite); err != nil {
			return err
		}
	}
	return nil
}

func DBEnable() bool {
	c := getDbConfig()
	return c.Enable
}

var tableModels []interface{}

// Migrate 初始化或升级表结构
func Migrate() error {
	DBPools.Range(func(k, db any) bool {
		if err := db.(*gorm.DB).AutoMigrate(tableModels...); err != nil {
			log.Fatalln("failed to migrate database:", k, ",err:", err)
			return false
		}
		return true
	})

	return nil
}

// RegTables 其他模块注册需要访问的表, 会被自动创建
func RegTables(tables ...interface{}) {
	tableModels = append(tableModels, tables...)
}

type InitDataHandler func() error

var _initHooks []InitDataHandler

// AddInitHook db连接后执行的函数， 可用于初始化数据等
func AddInitHook(h InitDataHandler) {
	_initHooks = append(_initHooks, h)
}

func InitData() error {
	for _, h := range _initHooks {
		if err := h(); err != nil {
			return err
		}
	}
	return nil
}
