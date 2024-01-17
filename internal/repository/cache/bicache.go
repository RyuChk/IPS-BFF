package cache

import (
	"context"

	"github.com/ZecretBone/ips-bff/internal/config"
	"github.com/allegro/bigcache/v3"
)

//go:generate mockgen -source=bicache.go -destination=mock_cache/mock_cache.go -package=mock_cache
type Service interface {
	Set(key string, entry []byte) error
	Get(key string) ([]byte, error)
}

type service struct {
	client *bigcache.BigCache
}

func ProvideCacheService(cfg config.CacheConfig) Service {
	config := bigcache.Config{
		Shards:             1024,
		CleanWindow:        cfg.CleanWindow,
		LifeWindow:         cfg.LifeWindow,
		MaxEntriesInWindow: cfg.HardMaxCacheSize,
	}

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		panic(err)
	}

	return &service{
		client: cache,
	}
}

func (s *service) Set(key string, entry []byte) error {
	return s.client.Set(key, entry)
}

func (s *service) Get(key string) ([]byte, error) {
	return s.client.Get(key)
}
