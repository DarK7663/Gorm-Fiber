package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=noonbyte password=noonbyte dbname=todo host=127.0.0.1 port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Task{})
}
