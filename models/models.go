package models

import (
	"fmt"
	"gin-template/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Model struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedOn  int  `json:"created_on"`
	ModifiedOn int  `json:"modified_on"`
	DeletedOn  int  `json:"deleted_on"`
}

var db *gorm.DB

func init() {
	var err error
	database := setting.Database
	args := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Host,
		database.Database)

	db, err = gorm.Open(database.Type, args)
	if err != nil {
		log.Printf("Failed to connect database: %s", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.Database.TablePrefix + defaultTableName
	}
	//defer db.Close()

	db.SingularTable(true)

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
}

func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if field, ok := scope.FieldByName("CreatedOn"); ok {
			if field.IsBlank {
				field.Set(nowTime)
			}
		}
		if field, ok := scope.FieldByName("ModifiedOn"); ok {
			if field.IsBlank {
				field.Set(nowTime)
			}
		}
	}
}

func updateTimestampForUpdateCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		if field, ok := scope.FieldByName("ModifiedOn"); ok {
			if field.IsBlank {
				field.Set(time.Now().Unix())
			}
		}
	}
}
