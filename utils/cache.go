package utils

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "1billion/models"
    "github.com/go-redis/redis/v8"
)

// CacheTTL defines the time-to-live (TTL) for cached data
const CacheTTL = 24 * time.Hour

// CacheKeyPrefix is the prefix for cache keys
const CacheKeyPrefix = "lead:"

// SetLeadCache sets lead information in the cache
func SetLeadCache(client *redis.Client, phoneNumber string, lead *models.Lead) error {
    key := fmt.Sprintf("%s%s", CacheKeyPrefix, phoneNumber)
    data, err := json.Marshal(lead)
    if err != nil {
        return err
    }
    return client.Set(context.Background(), key, data, CacheTTL).Err()
}

// GetLeadCache retrieves lead information from the cache
func GetLeadCache(client *redis.Client, phoneNumber string) (*models.Lead, error) {
    key := fmt.Sprintf("%s%s", CacheKeyPrefix, phoneNumber)
    data, err := client.Get(context.Background(), key).Bytes()
    if err != nil {
        if err == redis.Nil {
            return nil, nil // Cache miss
        }
        return nil, err
    }
    var lead models.Lead
    if err := json.Unmarshal(data, &lead); err != nil {
        return nil, err
    }
    return &lead, nil
}

// DeleteLeadCache deletes lead information from the cache
func DeleteLeadCache(client *redis.Client, phoneNumber string) error {
    key := fmt.Sprintf("%s%s", CacheKeyPrefix, phoneNumber)
    return client.Del(context.Background(), key).Err()
}
