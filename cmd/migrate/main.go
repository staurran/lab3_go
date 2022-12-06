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

	//err = db.AutoMigrate(&ds.Users{})
	if err != nil {
		panic("cant migrate db users")
	}

	err = db.AutoMigrate(&ds.Basket{})
	if err != nil {
		panic("cant migrate db basket")
	}

	err = db.AutoMigrate(&ds.Orders{})
	if err != nil {
		panic("cant migrate db orders")
	}

	err = db.AutoMigrate(&ds.Statuses{})
	if err != nil {
		panic("cant migrate db status")
	}

	err = db.AutoMigrate(&ds.GoodOrder{})
	if err != nil {
		panic("cant migrate db goodOrder")
	}
}
