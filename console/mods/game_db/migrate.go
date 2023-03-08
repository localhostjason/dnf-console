package game_db

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

func SyncDB() (err error) {
	if !DBEnable() {
		logrus.Info("no enable sql")
		return
	}

	err = Connect()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate:%v", err))
	}
	err = Migrate()
	if err != nil {
		return
	}

	err = InitData()
	return
}

func AutoMigrate() (err error) {
	if !DBEnable() {
		return
	}
	err = Connect()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate:%v", err))
	}
	return Migrate()
}
