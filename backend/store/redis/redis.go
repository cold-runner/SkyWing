package redis

import (
	"Skywing/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var RdbClient *Rdb

type Rdb struct {
	Client *redis.Client
}

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	c := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = c.Ping().Result()
	if err != nil {
		return err
	}
	RdbClient = &Rdb{Client: c}
	return nil
}

func Close() {
	_ = RdbClient.Client.Close()
}
