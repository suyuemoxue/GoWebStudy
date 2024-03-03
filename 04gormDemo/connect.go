package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DataBase *gorm.DB
var mysqllogger logger.Interface

type Student struct {
	Id     uint   `gorm:"size:10"`
	Name   string `gorm:"size:16"`
	Gender string `gorm:"size:6"`
	Age    int    `gorm:"size:3"`
}

var Studentlist []Student

func init() {
	username := "root"
	password := "sxj045115"
	host := "127.0.0.1"
	port := "3306"
	dbname := "gorm"
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, dbname)

	mysqllogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DataBase = db
}
