package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Product struct {
	gorm.Model
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
}

type SerialNumber struct {
	gorm.Model
	ID        int
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})
	if err != nil {
		panic(err)
	}

	////create catergory
	//category := Category{Name: "Cozinha"}
	//db.Create(&category)
	//
	////create product
	//db.Create(&Product{
	//	Name:       "Panela",
	//	Price:      99.0,
	//	CategoryID: 1,
	//})
	//
	////create a serial number
	//db.Create(&SerialNumber{
	//	Number:    "12345",
	//	ProductID: 1,
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}

	//var product []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&product)
	//for _, product := range product {
	//	println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}
}
