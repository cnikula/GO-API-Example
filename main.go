package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"your-project-path/controllers"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	r.GET("/products", controllers.GetProducts(db))
	r.GET("/products/:id", controllers.GetProductByID(db))
	r.DELETE("/products/:id", controllers.DeleteProduct(db))
	r.PUT("/products/:id", controllers.UpdateProduct(db))


	r.Run(":8080")
}

