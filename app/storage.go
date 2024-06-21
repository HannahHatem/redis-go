package main

import (
	// "fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"os"
	"strconv"
	"time"
)

type CacheValue struct {
	value         string
	startTime     time.Time
	span          time.Duration
	hasExpiration bool
}

var cache = make(map[string]CacheValue)

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

	cache[key] = CacheValue{
		value:         value,
		startTime:     time.Now(),
		span:          milliseconds,
		hasExpiration: expCheck,
	}

	return resp.WrapSimpleStringRESP("OK")
}

func GetMap(key string) string {
	value, exists := cache[key]

	if !exists {
		return resp.GetNullBulkStringRESP()
	}

	if !value.hasExpiration {
		return resp.WrapBulkStringRESP(value.value)
	}

	if time.Now().After(value.startTime.Add(value.span)) {
		return resp.GetNullBulkStringRESP()
	}

	return resp.WrapBulkStringRESP(value.value)
}
