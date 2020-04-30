package db

import (
	"fmt"
	"hub/src/app/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB

}

var MyDB *DB

func Setup(){
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		fmt.Println("Connect server error",err)
		return
	}
	MyDB = &DB{db}
	return
}

func GetMyDB() *DB{
	return MyDB
}

func (d *DB) GetGormDB() *gorm.DB{
	return MyDB.db
}

func (d *DB) Close(){
	MyDB.db.Close()
}




