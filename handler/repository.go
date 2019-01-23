package handler

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // ini harus ada
)

func getConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", os.Getenv("DBCONN"))
	if err != nil {
		return
	}
	return
}
