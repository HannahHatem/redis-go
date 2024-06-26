package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// var listen = flag.String("listen", ":6379", "listen address")
var port = flag.String("port", "6379", "address to listen to")
var replicaOf = flag.String("replicaof", "", "Replicate to another redis server")
var MasterReplIdValue = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
var MasterReplOffsetValue = "0"

func main() {
	fmt.Println("Logs from your program will appear here!")

	flag.Parse()
	sendHandshake()
	start()
}

func sendHandshake() {
	if *replicaOf != "" {
		replicaOfHostPort := strings.Split(*replicaOf, " ")
		replicaOfHost := replicaOfHostPort[0]
		replicaOfPort := replicaOfHostPort[1]

		conn, err := net.Dial("tcp", replicaOfHost+":"+replicaOfPort)
		if err != nil {
			log.Println("Failed to connect to replicaOf", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		sendPing := []string{"PING"}
		_, err = conn.Write([]byte(resp.WrapArrayRESP(sendPing)))
		if err != nil {
			log.Println("Failed to PING send handshake", err.Error())
			os.Exit(1)
		}

		// Sleep for 1 second to give the replicaOf server time to respond
		time.Sleep(1 * time.Second)
		replConf1 := []string{"REPLCONF", "listening-port", *port}
		_, err = conn.Write([]byte(resp.WrapArrayRESP(replConf1)))
		if err != nil {
			log.Println("Failed to REPLCONF1 send handshake", err.Error())
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		replConf2 := []string{"REPLCONF", "capa", "eof"}
		_, err = conn.Write([]byte(resp.WrapArrayRESP(replConf2)))
		if err != nil {
			log.Println("Failed to send handshake", err.Error())
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		psync := []string{"PSYNC", "?", "-1"}
		_, err = conn.Write([]byte(resp.WrapArrayRESP(psync)))
		if err != nil {
			log.Println("Failed to send handshake", err.Error())
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
	}
}

func start() {
	// l, err := net.Listen("tcp", *listen)
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		fmt.Println("Failed to bind to port 6379", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on port", fmt.Sprintf(":%s", *port))
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err.Error())
			os.Exit(1)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	// For each connection, we need to ensure that we close it when we're done.
	defer c.Close()
	for {
		buf := make([]byte, 128)
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection", err.Error())
			return
		}

		log.Printf("received from c.Read(buf): %s", buf)

		// Deserialize the byte array
		ans := resp.StartDeserializeParser(buf)
		if ans == nil {
			log.Println("Error deserializing byte array.")
			return
		} else if len(ans) == 0 {
			log.Println("Error,  byte array is empty.")
			return
		}

		// Master server
		cmdResult := handleCommand(ans)
		if cmdResult != "" {
			_, err = c.Write([]byte(cmdResult))
			if err != nil {
				fmt.Println("Error writing to connection c.Write()", err.Error())
				return
			}
		} else if cmdResult == "" {
			log.Println("Error, command result is empty.")
			return
		}

	}
}

func handleCommand(ans []string) string {
	command := strings.ToUpper(ans[0])

	switch command {
	case "PING":
		wrappedPong := resp.WrapSimpleStringRESP("PONG")
		return wrappedPong
	case "REPLCONF":
		wrappedReplConf := resp.WrapSimpleStringRESP("OK")
		return wrappedReplConf
	case "PSYNC":
		psyncResult := PsyncCommand(ans)
		return psyncResult
	case "SET":
		setResult := SetCommand(ans)
		return setResult
	case "GET":
		getResult := GetCommand(ans)
		return getResult
	case "ECHO":
		return resp.WrapBulkStringRESP(ans[1])
	case "INFO":
		info := InfoCommand(ans[1], *replicaOf)
		return info
	default:
		return resp.WrapBulkStringRESP("ERR unknown command")
	}
}
