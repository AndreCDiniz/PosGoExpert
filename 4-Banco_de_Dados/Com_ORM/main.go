package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	//create one product
	//db.Create(&Product{
	//	Name:  "PC Gamer",
	//	Price: 1000.00,
	//})

	//create batchn
	products := []Product{
		{Name: "PC Gamer", Price: 1000.00},
		{Name: "PC Normal", Price: 500.00},
		{Name: "PC Ruim", Price: 200.00},
	}
	db.Create(&products)
}
