package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/resp"
	// "strings"
)

var listen = flag.String("listen", ":6379", "listen address")

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	flag.Parse()
	start()
	// cmd := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	// byteArray := []byte(cmd)
	// fmt.Println("byteArray: ", byteArray)
	// resp.StartDeserializeParser(byteArray)
}

func start() {
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("Failed to bind to port 6379", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on port", *listen)
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
		} else if ans[0] == "SET" {
			setResult := SetMap(ans[1], ans[2])
			_, err = c.Write([]byte(setResult))
		} else if ans[0] == "GET" {
			getResult := GetMap(ans[1])
			_, err = c.Write([]byte(getResult))
		} else {
			_, err = c.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(ans[1]), ans[1])))
		}

		if err != nil {
			fmt.Println("Error writing to connection c.Write()", err.Error())
			return
		}
	}
}
