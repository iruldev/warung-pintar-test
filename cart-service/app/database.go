package app

import (
	"github.com/go-redis/redis"
	"github.com/iruldev/warung-pintar-test/cart-service/helper"
	"log"
	"os"
)

func NewDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	s := client.Ping()
	r, err := s.Result()
	helper.PanicIfError(err)

	log.Printf("Initializing a Redis Connection successful, with result : %s", r)

	return client
}
