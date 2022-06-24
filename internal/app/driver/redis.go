package driver

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheOption properties for cache DB
type CacheOption struct {
	IsEnable           bool
	Host               string
	Port               int
	Password           string
	Namespace          int
	DialConnectTimeout time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	MaxConnLifetime    time.Duration
}

// NewCache create cache client
func NewCache(option CacheOption) *redis.Client {
	redisOptions := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", option.Host, option.Port),
		DB:           option.Namespace,
		DialTimeout:  option.DialConnectTimeout,
		ReadTimeout:  option.ReadTimeout,
		WriteTimeout: option.WriteTimeout,
		PoolSize:     option.MaxActive,
		MinIdleConns: option.MaxIdle,
		MaxConnAge:   option.MaxConnLifetime,
	}

	if option.Password != "" {
		redisOptions.Password = option.Password
	}

	return redis.NewClient(redisOptions)
}
