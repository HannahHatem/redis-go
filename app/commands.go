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
