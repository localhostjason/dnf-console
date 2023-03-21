package game_db

import (
	"errors"
	"fmt"
	"github.com/localhostjason/webserver/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"sync"
)

// ConnectWithSqliteConfig 连接，检验配置是否正确
func ConnectWithSqliteConfig(cfgs []SqliteDBConfig) error {
	var eMsg = &sync.Map{}
	var wg = &sync.WaitGroup{}

	for _, c := range cfgs {
		wg.Add(1)

		go func(c SqliteDBConfig, eMsg *sync.Map) {
			defer wg.Done()
			var dbFile string
			if filepath.IsAbs(c.DbFile) {
				dbFile = c.DbFile
			} else {
				exePath, _ := os.Getwd()
				dbFile = filepath.Join(exePath, c.DbFile)
			}

			path, _ := filepath.Split(dbFile)
			if !util.PathExists(path) {
				_ = os.MkdirAll(path, os.ModePerm)
			}

			db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
				FullSaveAssociations:   true,
				SkipDefaultTransaction: true,
				NamingStrategy:         schema.NamingStrategy{SingularTable: true},
			})
			if err != nil {
				eMsg.Store("error", fmt.Sprintf("db file:%s, err:%s", dbFile, err))
				return
			}

			if c.Debug {
				db = db.Debug()
			}

			if err != nil {
				eMsg.Store("error", fmt.Sprintf("failed to connect databse %s:%v", c.DbFile, err))
				return
			}
			DBPools.Add(c.Key, db)
		}(c, eMsg)

	}

	wg.Wait()

	err, ok := eMsg.Load("error")
	if ok {
		return errors.New(err.(string))
	}
	return nil
}
