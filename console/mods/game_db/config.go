package game_db

import "github.com/localhostjason/webserver/server/config"

const _key = "game_db"

type MysqlDBConfig struct {
	Key             string `json:"key"`
	User            string `json:"user"`
	Password        string `json:"password"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	DB              string `json:"db"`
	Charset         string `json:"charset"`
	Timeout         int    `json:"timeout"`
	MultiStatements bool   `json:"multi_statements"`
	Debug           bool   `json:"debug"`
}

type SqliteDBConfig struct {
	Key    string `json:"key"`
	DbFile string `json:"db_file"`
	Debug  bool   `json:"debug"`
}

type DbConfig struct {
	Enable bool             `json:"enable"`
	Mysql  []MysqlDBConfig  `json:"mysql"`
	Sqlite []SqliteDBConfig `json:"sqlite"`
}

func init() {
	mc := MysqlDBConfig{
		Key:             "TestMysql1",
		User:            "root",
		Password:        "123456",
		Host:            "127.0.0.1",
		Port:            3306,
		DB:              "test",
		Charset:         "utf8mb4",
		MultiStatements: false,
		Timeout:         5,
		Debug:           false,
	}

	sc := SqliteDBConfig{
		Key:    "TestSqlite1",
		DbFile: "data/data.db",
		Debug:  false,
	}

	c := DbConfig{
		Enable: true,
		Mysql:  []MysqlDBConfig{mc},
		Sqlite: []SqliteDBConfig{sc},
	}
	_ = config.RegConfig(_key, c)
}

func getMysqlConfig() []MysqlDBConfig {
	var c DbConfig
	_ = config.GetConfig(_key, &c) // todo return error
	return c.Mysql
}

func getSqliteConfig() []SqliteDBConfig {
	var c DbConfig
	_ = config.GetConfig(_key, &c) // todo return error
	return c.Sqlite
}

func getDbConfig() DbConfig {
	var c DbConfig
	_ = config.GetConfig(_key, &c)
	return c
}
