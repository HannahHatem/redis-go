package main

import (
	// "fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/codecrafters-io/redis-starter-go/resp"
)

type cacheValue struct {
	value         string
	startTime     time.Time
	span          time.Duration
	hasExpiration bool
}

var redisStore = struct {
	sync.RWMutex
	cache map[string]cacheValue
}{cache: make(map[string]cacheValue)}

func SetMap(key string, value string, strSpan string) string {
	expCheck := false
	var milliseconds time.Duration

	if key == "" || value == "" {
		log.Println("Key or value is nil")
		os.Exit(1)
	} else if strSpan != "" {
		expCheck = true
		mills, err := strconv.Atoi(strSpan)
		if err != nil {
			log.Println("Error converting string to int")
			os.Exit(1)
		}
		milliseconds = time.Duration(mills) * time.Millisecond
	}

	redisStore.RWMutex.Lock()
	redisStore.cache[key] = cacheValue{
		value:         value,
		startTime:     time.Now(),
		span:          milliseconds,
		hasExpiration: expCheck,
	}
	redisStore.RWMutex.Unlock()

	return resp.WrapSimpleStringRESP("OK")
}

func GetMap(key string) string {
	redisStore.RWMutex.RLock()

	value, exists := redisStore.cache[key]
	redisStore.RWMutex.RUnlock()

	if !exists {
		return resp.GetNullBulkStringRESP()
	}

	if !value.hasExpiration {
		return resp.WrapBulkStringRESP(value.value)
	}

	if time.Now().After(value.startTime.Add(value.span)) {
		redisStore.RWMutex.Lock()
		delete(redisStore.cache, key)
		redisStore.RWMutex.Unlock()

		return resp.GetNullBulkStringRESP()
	}

	return resp.WrapBulkStringRESP(value.value)
}
