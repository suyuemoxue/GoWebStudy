package main

import (
	"gorm.io/gorm"
)

func main() {
	DataBase = DataBase.Session(&gorm.Session{Logger: mysqllogger})

	//where查询
	//DataBase.Where("name != ?", "墨雪2").Find(&studentlist)
	//fmt.Println(studentlist)

	//DataBase.Where("name like ?", "墨%").Find(&studentlist)
	//fmt.Println(studentlist)

	//DataBase.Where("age > 23 and gender = ?", "男").Find(&studentlist)
	//fmt.Println(studentlist)

	//DataBase.Where("age > 23").Or("age = 20").Find(&studentlist)
	//fmt.Println(studentlist)

	//排序
	//DataBase.Order("age desc").Where("age > 23").Find(&studentlist)
	//fmt.Println(studentlist)

	//分页
	//DataBase.Limit(3).Offset(0).Order("age desc").Where("age > 23").Find(&studentlist)
	//fmt.Println(studentlist)
	//DataBase.Limit(3).Offset(3).Order("age desc").Where("age > 23").Find(&studentlist)
	//fmt.Println(studentlist)

	//分组查询
	//var count []int
	//DataBase.Table("students").Select("count(id)").Group("gender").Scan(&count)
	//fmt.Println(count)
}
