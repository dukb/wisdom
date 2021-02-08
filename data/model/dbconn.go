package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wisdom/config"
)

func NewDBEngine(databasesetting *config.DatabaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databasesetting.User,
		databasesetting.Password,
		databasesetting.Host,
		databasesetting.DBname,
	))
	if err != nil {
		return nil, err
	}
	if databasesetting.LogMode == true {
		db.LogMode(true)
	}

	return db, nil
}
