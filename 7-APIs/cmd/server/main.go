package main

import (
	"github.com/AndreCDiniz/PosGoExpert/7-APIs/configs"
	"github.com/AndreCDiniz/PosGoExpert/7-APIs/internal/entity"
	"github.com/AndreCDiniz/PosGoExpert/7-APIs/internal/infra/database"
	"github.com/AndreCDiniz/PosGoExpert/7-APIs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("JwtExpiresIn", config.JwtExpiresIn))

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(config.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/", productHandler.GetProducts)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	log.Println("Server is running")
	http.ListenAndServe(":8000", router)

}
