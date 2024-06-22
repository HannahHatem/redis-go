package main

import (
	"log"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/resp"
)

func PsyncCommand(ans []string) string {
	log.Println("Psync command called")

	if ans[1] == "?" {
		fullResync := "FULLRESYNC " + MasterReplIdValue + " 0\r\n"
		return resp.WrapSimpleStringRESP(fullResync)
	}

	return resp.GetNullBulkStringRESP()
}

func InfoCommand(section string, replicaPort string) string {
	log.Println("Info command called with section: ", section)

	infoResultResp := ""
	masterReplID := "master_replid:" + MasterReplIdValue + "\r\n"
	masterReplOffset := "master_repl_offset:" + MasterReplOffsetValue + "\r\n"
	roleInfo := "role:master"

	if strings.ToLower(section) == "replication" {
		if replicaPort != "" {
			roleInfo = "role:slave"
		}
	}

	infoResultResp = roleInfo + "\r\n" + masterReplID + masterReplOffset

	return resp.WrapBulkStringRESP(infoResultResp)
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
