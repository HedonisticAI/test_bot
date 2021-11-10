package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Number  int
	ID      int
	Request string
	Adm     int
}

func inidb() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{Number: 1, ID: 288654334, Request: "I am admin lol", Adm: 1})

}

func addtodb(rs string, id int) {

}

func userhistory(ID int) string {

	var res string
	res = "aaaaa"
	return res
}
