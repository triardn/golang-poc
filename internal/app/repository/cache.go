package repository

import (
	"context"
	"time"
)

// ICacheRepo interface for cache repo
type ICacheRepo interface {
	WriteCache(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
	WriteCacheIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
	GetCache(ctx context.Context, key string) (data []byte, err error)
	DeleteCache(ctx context.Context, key string) (err error)
}

type cacheRepo struct {
	opt Option
}

// NewCacheRepository initiate cache repo
func NewCacheRepository(opt Option) ICacheRepo {
	return &cacheRepo{
		opt: opt,
	}
}

// WriteCache this will and must write the data to cache with corresponding key using locking
func (c *cacheRepo) WriteCache(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	return c.opt.CacheClient.SetEX(ctx, key, data, ttl).Err()
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteCacheIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	return c.opt.CacheClient.SetNX(ctx, key, data, ttl).Err()
}

func (c *cacheRepo) DeleteCache(ctx context.Context, key string) (err error) {
	return c.opt.CacheClient.Del(ctx, key).Err()
}

// GetCache Get cache as byte. Method who call this need to unmarshall the data
func (c *cacheRepo) GetCache(ctx context.Context, key string) (data []byte, err error) {
	return c.opt.CacheClient.Get(ctx, key).Bytes()
}
