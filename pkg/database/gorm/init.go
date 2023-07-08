package database

import (
	dbgorm "zombie-golang/pkg/database/dbgorm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Bootstrap() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&dbgorm.ArmorModel{})
	return db
}
