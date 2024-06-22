package main

import (
	"encoding/hex"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"strconv"
	"strings"
)

var emptyRDB, _ = hex.DecodeString("524544495330303131fa0972656469732d76657205372e322e30fa0a72656469732d62697473c040fa056374696d65c26d08bc65fa08757365642d6d656dc2b0c41000fa08616f662d62617365c000fff06e3bfec0ff5aa2")

func PsyncCommand(ans []string) string {
	log.Println("Psync command called")

	if ans[1] == "?" {
		fullResync := "FULLRESYNC " + MasterReplIdValue + " 0"
		return resp.WrapSimpleStringRESP(fullResync) + GetRDBFile()
	}
	return resp.GetNullBulkStringRESP()
}

func GetRDBFile() string {
	log.Println("Sending RDB file")

	s := string(emptyRDB)
	send := "$" + strconv.Itoa(len(s)) + "\r\n" + s

	return send
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
