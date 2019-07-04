package models

import (
	"gin-template/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"

	"log"
)

type Model struct {
	ID         uint `gorm:"primary_key"`
	CreatedOn  time.Time
	ModifiedOn time.Time
	DeletedOn  *time.Time
}

var db *gorm.DB

func init() {
	//dbConfig := setting.Database
	//connection := fmt.Sprintf("%s:%s@/%s", dbConfig.Username, dbConfig.Password, dbConfig.Database)
	//db, err := gorm.Open(dbConfig.Type, connection)

	var err error
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Failed to connect database: %s", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.Database.TablePrefix + defaultTableName
	}
	//defer db.Close()

	db.SingularTable(true)
}
