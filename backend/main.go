package main

import (
	"log"
	"os"
	"time"

	"github.com/Kapitar/url-shortener/handler"
	"github.com/Kapitar/url-shortener/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) { c.String(200, "ok") })

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortLink(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // local fallback
	}
	addr := "0.0.0.0:" + port
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server on %s: %v", addr, err)
	}
}
