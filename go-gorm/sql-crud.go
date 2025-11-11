package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id    int            `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	Name  sql.NullString // 姓名
	Age   int            // 年龄
	Grade sql.NullString // 年级
}

func main() {
	db, err := gorm.Open(mysql.Open("root:hycg8888@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Student{})

	// student := &Student{
	// 	Name:  sql.NullString{String: "Alice", Valid: true},
	// 	Age:   20,
	// 	Grade: sql.NullString{String: "A", Valid: true},
	// }
	// db.Create(student)

	// 插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	// db.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", "张三", 20, "三年级")

	// 查询 student 表中所有年龄大于 18 岁的学生信息
	var students []Student
	db.Raw("SELECT * FROM students WHERE age > ?", 18).Scan(&students)
	for _, student := range students {
		fmt.Println(student)
	}
	// 姓名为 "张三" 的学生年级更新为 "四年级"
	db.Exec("UPDATE students SET grade = ? WHERE name = ?", "四年级", "张三")

	// 删除 students 表中年龄小于 15 岁的学生记录
	db.Exec("DELETE FROM students WHERE age < ?", 15)

}
