package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/resp"
)

// var listen = flag.String("listen", ":6379", "listen address")
var port = flag.String("port", "6379", "address to listen to")

func main() {
	fmt.Println("Logs from your program will appear here!")

	flag.Parse()
	start()
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

		if ans[0] == "PING" {
			wrapedPong := resp.WrapSimpleStringRESP("PONG")
			_, err = c.Write([]byte(wrapedPong))
		} else if strings.ToUpper(ans[0]) == "SET" {
			setResult := ""
			if len(ans) < 4 {
				setResult = SetMap(ans[1], ans[2], "")
			} else {
				setResult = SetMap(ans[1], ans[2], ans[4])
			}
			_, err = c.Write([]byte(setResult))
		} else if strings.ToUpper(ans[0]) == "GET" {
			getResult := GetMap(ans[1])
			_, err = c.Write([]byte(getResult))
		} else if strings.ToUpper(ans[0]) == "ECHO" {
			_, err = c.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(ans[1]), ans[1])))
		}

		if err != nil {
			fmt.Println("Error writing to connection c.Write()", err.Error())
			return
		}
	}
}
