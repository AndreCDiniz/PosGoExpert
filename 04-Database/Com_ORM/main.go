package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
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
	//products := []Product{
	//	{Name: "PC Gamer", Price: 1000.00},
	//	{Name: "PC Normal", Price: 500.00},
	//	{Name: "PC Ruim", Price: 200.00},
	//}
	//db.Create(&products)

	//selecionar o primeiro que tenha PC Gamer
	//var product Product
	//db.First(&product, "name = ?", "PC Gamer")
	//fmt.Println(product)

	//select all
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//O limit usamos para limitar a quantidade de coisas que buscamos
	//enquanto o offset é utilizado para paginação.
	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//usando where
	//var products []Product
	//db.Where("price > ?", 300).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//	//usando like
	//	//var products []Product
	//	//db.Where("name LIKE ?", "%Gamer").Find(&products)
	//	//for _, product := range products {
	//	//	fmt.Println(product)
	//}

	//Fazendo alteração do dado
	//var p Product
	//db.First(&p, 1)
	//p.Name = "Mouse"
	//db.Save(&p)

	//Fazendo deleção de um dado
	//var p2 Product
	//db.First(&p2, 1)
	//db.Delete(&p2)

}
