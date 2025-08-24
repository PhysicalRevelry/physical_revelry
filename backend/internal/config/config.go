package config

import (
    "os"
)

type Config struct {
    Port            string
    DatabaseURL     string
    Auth0Domain     string
    Auth0Audience   string
    AlgoliaAppID    string
    AlgoliaAPIKey   string
    AlgoliaIndex    string
    AWSRegion       string
    DynamoDBTable   string
}

func New() *Config {
    return &Config{
        Port:            getEnv("PORT", "8080"),
        DatabaseURL:     getEnv("DATABASE_URL", ""),
        Auth0Domain:     getEnv("AUTH0_DOMAIN", ""),
        Auth0Audience:   getEnv("AUTH0_AUDIENCE", ""),
        AlgoliaAppID:    getEnv("ALGOLIA_APP_ID", ""),
        AlgoliaAPIKey:   getEnv("ALGOLIA_API_KEY", ""),
        AlgoliaIndex:    getEnv("ALGOLIA_INDEX", "posts"),
        AWSRegion:       getEnv("AWS_REGION", "us-east-1"),
        DynamoDBTable:   getEnv("DYNAMODB_TABLE", "my-app-table"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}