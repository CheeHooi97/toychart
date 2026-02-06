package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Env                string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	SystemAesKey       string
	OSSEndpoint        string
	OSSAccessKeyID     string
	OSSAccessKeySecret string
	OSSBucket          string
	EbayClientId       string
	EbayClientSecret   string
)

// LoadConfig
func LoadConfig() {
	_ = godotenv.Load()

	Env = GetEnv("ENV")
	DBHost = GetEnv("POSTGRES_HOST")
	DBPort = GetEnv("POSTGRES_PORT")
	DBUser = GetEnv("POSTGRES_USER")
	DBPassword = GetEnv("POSTGRES_PASSWORD")
	DBName = GetEnv("POSTGRES_DATABASE")
	SystemAesKey = GetEnv("SYSTEM_AES_KEY")
	OSSEndpoint = GetEnv("OSS_ENDPOINT")
	OSSAccessKeyID = GetEnv("OSS_ACCESS_KEY_ID")
	OSSAccessKeySecret = GetEnv("OSS_ACCESS_KEY_SECRET")
	OSSBucket = GetEnv("OSS_BUCKET")
	EbayClientId = GetEnv("EBAY_CLIENT_ID")
	EbayClientSecret = GetEnv("EBAY_CLIENT_SECRET")
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}
