package main

import (
	// "fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"os"
	"time"
)

type CacheValue struct {
	value     string
	startTime time.Time
	span      time.Duration
}

var cache = make(map[string]string)

func SetMap(key string, value string) string {

	if key == "" || value == "" {
		log.Println("Key or value is nil")
		os.Exit(1)
	}

	cache[key] = value

	return resp.WrapSimpleStringRESP("OK")
}

func GetMap(key string) string {
	value, exists := cache[key]

	if !exists {
		return resp.GetNullBulkStringRESP()
	}

	return resp.WrapBulkStringRESP(value)
}
