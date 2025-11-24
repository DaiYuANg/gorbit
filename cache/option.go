package cache

import (
	"time"

	"github.com/eko/gocache/lib/v4/cache"
)

// CacheOption 功能型 Option
type CacheOption func(*cacheOptions[any])

type cacheOptions[T any] struct {
	DefaultTTL time.Duration
	Store      cache.Cache[T]
}

// 默认配置
//func defaultCacheOptions[T any]() *cacheOptions[T] {
//	return &cacheOptions[T]{
//		DefaultTTL: 5 * time.Minute,
//		Store:      memory.NewStore(),
//	}
//}
//
//// Option helpers
//func WithDefaultTTL(ttl time.Duration) CacheOption {
//	return func(o *cacheOptions) { o.DefaultTTL = ttl }
//}
//
//func WithMemoryStore() CacheOption {
//	return func(o *cacheOptions) { o.Store = memory.NewStore() }
//}
//
//func WithRedisStore(client *redis.Client, prefix string) CacheOption {
//	return func(o *cacheOptions) {
//		o.Store = redis.New(client, prefix)
//	}
//}
//
//func WithSerializer(serializer gocache.Serializer) CacheOption {
//	return func(o *cacheOptions) { o.Serializer = serializer }
//}
