package infrastructure

import (
	"phenikaa/utils"
	"strconv"

	"github.com/go-redis/redis"
)

func connectRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// InitRedis open connection to redis server
func InitRedis() error {
	var err error
	redisClient, err = connectRedis()
	if err != nil {
		return err
	}

	return nil
}

// DeleteAuth delete pair key value has key same givenUUID
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := redisClient.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func ClearAuth(userID uint) {
	list, _ := redisClient.Keys(utils.PatternGet(userID)).Result()
	for _, key := range list {
		redisClient.Del(key)
	}
}

func FetchAuth(accessUUID string) (uint, error) {
	userIDStr, err := redisClient.Get(accessUUID).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	return uint(userID), nil
}
