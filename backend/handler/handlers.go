package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Kapitar/url-shortener/shortener"
	"github.com/Kapitar/url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	UserId      string `json:"user_id" binding:"required,uuid"`
}

func CreateShortLink(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.OriginalUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.OriginalUrl, creationRequest.UserId)
	host := os.Getenv("BASE_URL")
	c.JSON(200, gin.H{
		"success":   true,
		"short_url": host + "/" + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	fmt.Println(shortUrl)
	initialUrl := store.GetOriginalUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
