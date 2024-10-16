package cache

import "time"

type CacheItem struct {
	value  any
	expiry time.Time
}
