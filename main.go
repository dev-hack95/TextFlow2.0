package main

import (
	"github/dev-hack95/Textflow/handlers"
	"github/dev-hack95/Textflow/utilities"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	utilities.EnableSQLDatabasesConfiguration()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	// Users
	router.POST("/v1/signup", handlers.SignUp)
	router.POST("/v1/signin", handlers.SignIn)
	// Upload Video to Minio s3 storage
	router.POST("/v1/upload", handlers.Upload)
	router.POST("/v1/producer", handlers.Produce)

	router.Run("0.0.0.0:8000")
}
