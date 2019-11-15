package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		for {
			for _, v := range `-\|/` {
				fmt.Printf("\r%cwaiting for connect...", v)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("\none client[%s] connect successfully!\n", conn.RemoteAddr())
		handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Fatal(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
