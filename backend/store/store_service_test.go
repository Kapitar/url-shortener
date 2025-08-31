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

	SaveUrlMapping(shortUrl, originalUrl)

	retrievedUrl := GetOriginalUrl(shortUrl)

	assert.Equal(t, originalUrl, retrievedUrl)
}
