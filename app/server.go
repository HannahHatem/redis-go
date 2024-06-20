package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"log"
	"net"
	"os"
)

var listen = flag.String("listen", ":6379", "listen address")

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	flag.Parse()
	// start()
	// cmd := "*2\r\n$4\r\nECHO\r\n$9\r\nblueberry\r\n"
	cmd := "+PONG\r\n"
	byteArray := []byte(cmd)
	fmt.Println("byteArray: ", byteArray)
	ans := resp.StartDeserializeParser(byteArray)
	fmt.Println(ans)
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

		_, err = c.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing to connection c.Write()", err.Error())
			return
		}
	}
}
