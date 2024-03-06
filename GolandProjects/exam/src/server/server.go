package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Received: %s", string(buf[:n]))
		data, _ := strconv.Atoi(string(buf[:n]))
		data *= 2
		_, err = conn.Write([]byte(strconv.Itoa(data) + "\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "185.139.70.64:8000")
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listening on port 8000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
