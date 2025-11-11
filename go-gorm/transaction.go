package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	Id      int             `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	Balance sql.NullFloat64 // 账户余额，允许为空
}

type Transaction struct {
	Id            int             `gorm:"primaryKey,autoIncrement"` // 主键，自动增长
	FromAccountId sql.NullInt64   // 转出账户ID，允许为空
	ToAccountId   sql.NullInt64   // 转入账户ID，允许为空
	Amount        sql.NullFloat64 // 转账金额，允许为空
}

func main() {
	db, err := gorm.Open(mysql.Open("root:hycg8888@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	// 并在 transaction 表中记录该笔转账信息。如果余额不足，则回滚事务
	err = db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		if err := tx.First(&fromAccount, 1).Error; err != nil {
			return err
		}
		if err := tx.First(&toAccount, 2).Error; err != nil {
			return err
		}
		transferAmount := 100.0

		if fromAccount.Balance.Valid && fromAccount.Balance.Float64 >= transferAmount {
			fromAccount.Balance.Float64 -= transferAmount
			toAccount.Balance.Float64 += transferAmount
			if err := tx.Save(&fromAccount).Error; err != nil {
				return err
			}
			if err := tx.Save(&toAccount).Error; err != nil {
				return err
			}
			transaction := Transaction{
				FromAccountId: sql.NullInt64{Int64: int64(fromAccount.Id), Valid: true},
				ToAccountId:   sql.NullInt64{Int64: int64(toAccount.Id), Valid: true},
				Amount:        sql.NullFloat64{Float64: transferAmount, Valid: true},
			}
			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
			return nil
		} else {
			return fmt.Errorf("insufficient funds in account %d", fromAccount.Id)
		}
	})
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction succeeded")
	}

}
