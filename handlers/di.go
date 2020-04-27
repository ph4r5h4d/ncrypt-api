package handlers

import (
	"github.com/go-redis/redis/v7"
	keyManager "ncrypt-api/key-manager"
	"ncrypt-api/models"
)

type DI struct {
	RedisClient *redis.Client
	models.Key
}

func BuildDI(config models.Config) (DI, error) {
	di := DI{}

	rc, err := buildRedisConnection(config.RedisConfig)
	if err != nil {
		return DI{}, err
	}
	di.RedisClient = rc

	privateKey, err := keyManager.LoadCurrentKey()
	if err != nil {
		return DI{}, err
	}
	di.Key.PrivateKey = privateKey
	di.Key.PublicKey = &privateKey.PublicKey
	di.Key.Id = keyManager.GetKeyId()

	return di, nil
}

func buildRedisConnection(config models.RedisConfig) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Database, // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		return nil, err
	}

	return c, nil
}
