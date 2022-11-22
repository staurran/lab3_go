package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab3/internal/app/ds"
	"lab3/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&ds.Goods{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Users{})
	if err != nil {
		panic("cant migrate db users")
	}

	err = db.AutoMigrate(&ds.Bucket{})
	if err != nil {
		panic("cant migrate db bucket")
	}

}
