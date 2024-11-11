package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	CategoryID uint
	Category   Category
}

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

	//create catergory
	//category := Category{Name: "Eletrônicos"}
	//db.Create(&category)

	//create product
	//db.Create(&Product{
	//	Name:       "Smartphone",
	//	Price:      2000,
	//	CategoryID: category.ID,
	//})

	var product []Product
	db.Preload("Category").Find(&product)
	for _, product := range product {
		println(product.Name, product.Category.Name)
	}
}
