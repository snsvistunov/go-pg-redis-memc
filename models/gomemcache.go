package models

import (
	"time"

	gc "github.com/patrickmn/go-cache"
)

const cacheTTL int = 15

var MC *gc.Cache

func ConnectMemcache() {
	MC = gc.New(time.Duration(cacheTTL)*time.Second, time.Duration(cacheTTL)*time.Second)

}

func SetMemcacheValue(c *gc.Cache, key string, value interface{}) {
	c.Set(key, value, gc.DefaultExpiration)
}

func GetMemcacheValue(c *gc.Cache, key string) (interface{}, bool) {
	x, found := c.Get(key)
	if found {
		return x, found
	}
	return nil, found
}
