package config

import (
	"fmt"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/redis/go-redis/v9"
)

func NewRedis(conf env.Cache) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Address, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	})

	return rdb
}
