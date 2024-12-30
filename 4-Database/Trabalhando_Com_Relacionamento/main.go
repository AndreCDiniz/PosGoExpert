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
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	//SerialNumber SerialNumber
}

//type SerialNumber struct {
//	gorm.Model
//	ID        int
//	Number    string
//	ProductID int
//}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Category{}, &Product{})
	if err != nil {
		panic(err)
	}

	////create catergory
	//category := Category{Name: "Cozinha"}
	//db.Create(&category)
	//
	//category2 := Category{Name: "Eletronico"}
	//db.Create(&category2)
	//
	////create product
	//db.Create(&Product{
	//	Name:       "Panela",
	//	Price:      99.0,
	//	Categories: []Category{category, category2},
	//})

	////create a serial number
	//db.Create(&SerialNumber{
	//	Number:    "12345",
	//	ProductID: 1,
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}

	//var product []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&product)
	//for _, product := range product {
	//	println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}
}
