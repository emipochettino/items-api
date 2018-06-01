package dao

import (
	"fmt"

	"github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	UserLocalMysql     = "root"
	PassLocalMysql     = ""
	DataBaseLocalMysql = "test"
)

var CategoryDAO CategoryDao

func init() {
	initializeDataBase()
	initializeDao()
}

func getConnection() (con *gorm.DB, err error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", UserLocalMysql, PassLocalMysql, DataBaseLocalMysql))
	if err != nil {
		logger.Error("Something went wrong when it tried to get the db connection")
		return nil, err
	}

	db.DB().SetMaxIdleConns(25)
	db.DB().SetMaxOpenConns(50)
	db.SingularTable(true)

	return db, err
}

func initializeDataBase() {
	db, err := getConnection()
	if err != nil {
		logger.Error("Something went wrong when it tried to get the db connection")
	}
	defer db.Close()

	db.DropTableIfExists(&entities.Category{})
	db.AutoMigrate(&entities.Category{})
}

func initializeDao() {
	CategoryDAO = CategoryDao{}
}
