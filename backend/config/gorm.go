package config

import (
	"backend/httpserver/models"
	"os"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("PGHOST") +
		" user=" + os.Getenv("PGUSER") +
		" password=" + os.Getenv("PGPASSWORD") +
		" dbname=" + os.Getenv("PGDATABASE") +
		" port=" + os.Getenv("PGPORT") +
		" sslmode=disable TimeZone=Asia/Jakarta"

	var db *gorm.DB
	var err error

	for i := 0; i < 3; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(time.Second * 5) // Tunggu 5 detik sebelum mencoba lagi
	}

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.UserModel{})
	return db, err
}
