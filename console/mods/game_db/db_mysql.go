// Package game_db Package db 处理数据库连接信息，实现Model对象的存储读取
package game_db

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectWithMysqlConfig 连接，检验配置是否正确
func ConnectWithMysqlConfig(cfgs []MysqlDBConfig) error {
	for _, c := range cfgs {
		multiStatements := "false"
		if c.MultiStatements {
			multiStatements = "true"
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=%s&charset=%s&parseTime=True&loc=Local&timeout=%ds&allowAllFiles=true",
			c.User, c.Password, c.Host, c.Port, c.DB, multiStatements, c.Charset, c.Timeout)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			FullSaveAssociations:   true,
			SkipDefaultTransaction: true,
			NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		})

		if c.Debug {
			db = db.Debug()
		}

		if err != nil {
			return errors.New(fmt.Sprintf("failed to connect databse %s:%v", dsn, err))
		}

		DBPools.Add(c.Key, db)
	}

	return nil
}

// Close todo 关闭db连接
func Close() error {
	return nil
}
