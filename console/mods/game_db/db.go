package game_db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DBPools = new(DBPool)

func Connect() error {
	c := getDbConfig()
	if !c.Enable {
		return nil
	}

	if err := ConnectWithMysqlConfig(c.Mysql); err != nil {
		return err
	}

	if err := ConnectWithSqliteConfig(c.Sqlite); err != nil {
		return err
	}
	return nil
}

func DBEnable() bool {
	c := getDbConfig()
	return c.Enable
}

var tableModels = make(map[any][]interface{})

// Migrate 初始化或升级表结构
func Migrate() error {
	DBPools.Range(func(k, db any) bool {
		if err := db.(*gorm.DB).AutoMigrate(tableModels[k]...); err != nil {
			log.Fatalln("failed to migrate database:", k, ",err:", err)
			return false
		}
		return true
	})

	return nil
}

// RegTables 其他模块注册需要访问的表, 会被自动创建
func RegTables(k any, tables ...interface{}) {
	tableModels[k] = append(tableModels[k], tables...)
}

type InitGameDataHandler func() error

var _initGameHooks []InitGameDataHandler

// AddInitHook db连接后执行的函数， 可用于初始化数据等
func AddInitHook(h InitGameDataHandler) {
	_initGameHooks = append(_initGameHooks, h)
}

func InitData() error {
	for _, h := range _initGameHooks {
		if err := h(); err != nil {
			return err
		}
	}
	return nil
}
