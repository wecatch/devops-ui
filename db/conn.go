package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// init mysql init
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

// DB for database
var DB *gorm.DB

// NewDB init db
func NewDB(
	databaseAddress, databaseUser, databasePasswd, databaseName string, debug bool) {
	var err error
	logger = log.Logger("db")
	DB, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		databaseUser, databasePasswd, databaseAddress, databaseName))
	if err != nil {
		logger.Fatal(err)
	}

	DB.LogMode(debug)
	DB.SingularTable(true)
	// check db conn
	DB.Raw("select 1")
}
