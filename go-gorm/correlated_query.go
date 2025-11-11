/*
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表
*/
package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id           int            `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	Username     sql.NullString // 用户名，允许为空
	Email        sql.NullString // 邮箱，允许为空
	Posts        []Post         `gorm:"foreignKey:UserId"` // 一对多关系：一个用户可以有多篇文章
	ArticleCount int            // 文章数量统计字段
}
type Post struct {
	Id       int            `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	Title    sql.NullString // 文章标题，允许为空
	Content  sql.NullString // 文章内容，允许为空
	UserId   sql.NullInt64  // 外键，关联 User 表的 Id，允许为空
	Comments []Comment      `gorm:"foreignKey:PostId"` // 一对多关系：一篇文章可以有多条评论
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.First(&user, p.UserId.Int64).Error; err != nil {
		return err
	}
	// 假设 User 模型有一个 ArticleCount 字段用于统计文章数量
	user.ArticleCount++
	return tx.Save(&user).Error
}

type Comment struct {
	Id      int            `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	Content sql.NullString // 评论内容，允许为空
	PostId  sql.NullInt64  // 外键，关联 Post 表的 Id，允许为空
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.First(&post, c.PostId.Int64).Error; err != nil {
		return err
	}
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", post.Id).Count(&commentCount).Error; err != nil {
		return err
	}
	if commentCount == 0 {
		// 假设 Post 模型有一个 CommentStatus 字段用于表示评论状态
		// post.CommentStatus = "无评论"
		return tx.Save(&post).Error
	}
	return nil
}

func main() {
	db, err := gorm.Open(mysql.Open("root:hycg8888@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&User{})

	// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	var user User
	db.Preload("Posts.Comments").First(&user, 1) // 假设查询ID为1的用户
	fmt.Println(user)
	// 使用Gorm查询评论数量最多的文章信息
	var post Post
	db.Preload("Comments").Order("LENGTH(Comments) DESC").First(&post)
	fmt.Println(post)

}
