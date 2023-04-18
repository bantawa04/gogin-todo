package db

import (
	"fmt"
	"todo-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "127.0.0.1", "3306", "todo"))
	url := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Todo{})
	return db
}
