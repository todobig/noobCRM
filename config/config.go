package config

import "os"

// Config represents application configuration
type Config struct {
    MongoDBURI  string
    MongoDBName string
    RedisAddr   string
    // Add more configuration fields as needed
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
    return &Config{
        MongoDBURI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
        MongoDBName: getEnv("MONGODB_NAME", "your_database_name"),
        RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
        // Initialize other configuration fields here
    }
}

// getEnv retrieves the value of an environment variable or returns a default value if not set
func getEnv(key, fallbackValue string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallbackValue
}
