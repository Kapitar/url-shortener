package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndGet(t *testing.T) {
	originalUrl := "https://google.com"
	shortUrl := "test"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"

	SaveUrlMapping(shortUrl, originalUrl, userUUId)

	retrievedUrl := GetOriginalUrl(shortUrl)

	assert.Equal(t, originalUrl, retrievedUrl)
}
