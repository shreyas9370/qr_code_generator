package main

import (
	"qr-code-generator/controllers"
	"qr-code-generator/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Init()

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/generate", controllers.GenerateQRCode)
	router.GET("/qr/:id", controllers.RedirectTOURL)

	router.Run(":8080")
}
