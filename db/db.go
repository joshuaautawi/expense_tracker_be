package db

import (
	"github.com/joshua/expense_tracker/expense"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/expense_tracker"
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&expense.Expense{})
	return DB
}
