package main

import (
	"fmt"
	"gorm.io/gorm"
)

// 添加记录、查询记录
func main() {
	DataBase = DataBase.Session(&gorm.Session{Logger: mysqllogger})
	//DataBase.AutoMigrate(&Student{})

	//var studentlist []Student
	//原生SQL示例
	DataBase.Raw("select * from students where age >= 24").Scan(&Studentlist)
	fmt.Println(Studentlist)

	//添加记录
	//for i := 0; i < 10; i++ {
	//	studentlist = append(studentlist, Student{
	//		Name:   fmt.Sprintf("墨雪%d", i+1),
	//		Gender: "男",
	//		Age:    18 + i + 1,
	//	})
	//}
	//
	//err := DataBase.Create(&studentlist).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//var student Student
	//单条记录的查询
	//DataBase.Take(&student)
	//fmt.Println(student)

	//根据ID查找
	//DataBase.Take(&student, "3")
	//fmt.Println(student)

	//根据name查找
	//DataBase.Take(&student, "name = ?", "墨雪5")
	//fmt.Println(student)

	//查询多条记录
	//cout := DataBase.Find(&studentlist).RowsAffected //统计有多少条记录
	//fmt.Println("共有", cout, "条数据")
	//for _, s := range studentlist {
	//	fmt.Println(s)
	//}

	//根据主键ID多条查询
	//DataBase.Find(&studentlist, []int{2, 4, 6})
	//fmt.Println(studentlist)
	//根据name多条查询
	//DataBase.Find(&studentlist, "name in (?)", []string{"墨雪3", "墨雪6"})
	//fmt.Println(studentlist)

	// 更新  用save
	//DataBase.Take(&student, 6)
	//student.Age = 18
	//DataBase.Save(&student)

	// 更新  用update
	//DataBase.Find(&studentlist, []int{3, 4, 5}).Update("gender", "女")

	// 删除
	//DataBase.Find(&studentlist, []int{9, 10})
	//DataBase.Delete(studentlist)
}
