package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	// init mysql init
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

type DatabaseConf struct {
	DatabaseAddress, DatabaseUser, DatabasePasswd, DatabaseName string
	Debug                                                       bool
}

type DatabaseConfMap map[string]DatabaseConf

var dbconfigMap DatabaseConfMap

// NewDB init db
func NewDB(
	databaseAddress, databaseUser, databasePasswd, databaseName string, debug bool) {
	logger = log.Logger("db")
	if dbconfigMap == nil {
		dbconfigMap = make(DatabaseConfMap)
	}
	dbconfig := DatabaseConf{}
	dbconfig.DatabaseAddress = databaseAddress
	dbconfig.DatabaseUser = databaseUser
	dbconfig.DatabasePasswd = databasePasswd
	dbconfig.DatabaseName = databaseName
	dbconfig.Debug = debug
	dbconfigMap[databaseName] = dbconfig
}

func MakeNewDB(dbName string) (*gorm.DB, error) {
	dbconfig, ok := dbconfigMap[dbName]
	if !ok {
		return nil, errors.New("dbName config not exist")
	}
	logger := log.Logger("makeNewDB")
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.DatabaseUser, dbconfig.DatabasePasswd, dbconfig.DatabaseAddress, dbconfig.DatabaseName))
	if err != nil {
		logger.Fatal(err)
	}

	db.LogMode(dbconfig.Debug)
	db.SingularTable(true)
	// check db conn
	db.Raw("select 1")

	return db, nil
}

func Session() *gorm.DB {
	db, err := MakeNewDB("devops")
	if err != nil {
		return nil
	}

	return db
}
