package game_db

import (
	"errors"
	"fmt"
	"github.com/localhostjason/webserver/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// ConnectWithSqliteConfig 连接，检验配置是否正确
func ConnectWithSqliteConfig(cfgs []SqliteDBConfig) error {
	var wg = &sync.WaitGroup{}

	for _, c := range cfgs {
		wg.Add(1)

		go func(c SqliteDBConfig) {
			defer wg.Done()
			err := connectSqliteOne(c)
			if err != nil {
				log.Fatal(err)
			}
		}(c)

	}

	wg.Wait()

	if err := Migrate(); err != nil {
		return err
	}

	if err := InitData(); err != nil {
		return err
	}

	return nil
}

func connectSqliteOne(c SqliteDBConfig) error {
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
		return errors.New(fmt.Sprintf("db file:%s, err:%s", dbFile, err))
	}

	if c.Debug {
		db = db.Debug()
	}

	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect databse %s:%v", c.DbFile, err))
	}
	DBPools.Add(c.Key, db)
	return nil
}
