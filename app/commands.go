package main

import (
	"log"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/resp"
)

func InfoCommand(section string, replicaPort string) string {
	log.Println("Info command called with section: ", section)
	if strings.ToLower(section) == "replication" {
		roleInfo := "role:master"
		if replicaPort != "" {
			roleInfo = "role:slave"
		}
		return resp.WrapBulkStringRESP(roleInfo)
	}
	return resp.GetNullBulkStringRESP()
}

func SetCommand(ans []string) string {
	setResult := ""
	if len(ans) < 4 {
		setResult = SetMap(ans[1], ans[2], "")
	} else {
		setResult = SetMap(ans[1], ans[2], ans[4])
	}
	return setResult
}

func GetCommand(ans []string) string {
	getResult := GetMap(ans[1])
	return getResult
}
