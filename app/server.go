package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"net"
	"os"
	"strings"
)

// var listen = flag.String("listen", ":6379", "listen address")
var port = flag.String("port", "6379", "address to listen to")
var replicaOf = flag.String("replicaof", "", "Replicate to another redis server")

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
		// _, err = conn.Write([]byte(fmt.Sprintf("REPLCONF listening-port %s\r\n", *port)))
		// if err != nil {
		// 	log.Println("Failed to send handshake", err.Error())
		// 	os.Exit(1)
		// }

		sendPing := []string{"PING"}
		_, err = conn.Write([]byte(resp.WrapArrayRESP(sendPing)))
		if err != nil {
			log.Println("Failed to send handshake", err.Error())
			os.Exit(1)
		}
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
