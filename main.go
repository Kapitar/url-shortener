package main

import (
	"fmt"
	"log"

	"github.com/Kapitar/url-shortener/handler"
	"github.com/Kapitar/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortLink(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
