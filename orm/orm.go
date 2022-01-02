package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ToDo struct {
	Id          uint
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

var DSN = "root:nikhil@/todo"

func GetDatabase() *gorm.DB {
	DB, err := gorm.Open(mysql.New(mysql.Config{DSN: DSN}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&ToDo{})
	return DB
}
