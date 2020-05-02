package models

import (
	"fmt"
	myDB "hub/src/app/db"
	"hub/src/app/pkg/setting"
	"regexp"
	"strconv"
	"testing"
)

func Test_AutoMigrate(t *testing.T) {
	setting.Setup()
	//db1 := db.NewClient("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	//db1 := db.NewClient("tophub:FEMGhS36LMPThsrh@tcp(47.97.98.245:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	myDB.Setup()
	myDB.GetMyDB().GetGormDB().AutoMigrate(&Item{})
	fmt.Println(myDB.GetMyDB())

}

func Test_Re(t *testing.T){
	reg, _ := regexp.Compile("\\d+")
	comNum2 := reg.Find([]byte("10万+"))
	comNum3, _ := strconv.Atoi(string(comNum2))
	t.Log(comNum3)
}