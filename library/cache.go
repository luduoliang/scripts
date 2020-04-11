package library

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var MemCache *cache.Cache

func init() {
	//内存缓存
	MemCache  = cache.New(cache.NoExpiration, 86400*time.Second)
}