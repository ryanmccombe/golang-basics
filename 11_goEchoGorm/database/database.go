package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Settings struct {
	Connection *gorm.DB
}

func Connect() *Settings {
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	settings := &Settings{Connection: db}

	return settings
}
