package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var listen = flag.String("listen", ":6379", "listen address")

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	flag.Parse()

	err := run()
	if err != nil {
		fmt.Println("Error running server", err.Error())
		os.Exit(1)
	}
}

func run() error {
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("Failed to bind to port 6379", err.Error())
		return errors.New("failed to bind to port 6379")
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection", err.Error())
		return errors.New("error accepting connection")
	}

	defer c.Close()

	buf := make([]byte, 128)
	_, err = c.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection", err.Error())
		return errors.New("error reading from connection")
	}

	log.Printf("received from c.Read(buf): %s", buf)

	_, err = c.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Error writing to connection c.Write()", err.Error())
		return errors.New("error writing to connection")
	}

	return nil
}
