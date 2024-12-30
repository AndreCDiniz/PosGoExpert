package database

import (
	"fmt"
	"github.com/AndreCDiniz/PosGoExpert/7-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func setupDb(t *testing.T) (*gorm.DB, *Product) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})
	return db, NewProduct(db)
}

func TestCreateNewProduct(t *testing.T) {
	_, productDB := setupDb(t)

	product, _ := entity.NewProduct("Product 1", 10.0)
	err := productDB.Create(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFinadAllProducts(t *testing.T) {
	db, productDB := setupDb(t)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64())
		assert.NoError(t, err)
		db.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductById(t *testing.T) {
	db, productDB := setupDb(t)

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	db.Create(product)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", productFound.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, productDB := setupDb(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	db.Create(product)

	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", productFound.Name)
}

func TestDeleteProduct(t *testing.T) {
	_, productDB := setupDb(t)

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.ErrorContains(t, err, "record not found")
}
