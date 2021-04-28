package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Falha ao conectar com o banco de dados")
	}

	db.AutoMigrate(&Book{})

	return db
}
