package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var isClient bool

func init() {
	flag.BoolVar(&isClient, "b", false, "")
}

func main() {
	flag.Parse()

	if isClient {
		conn, err := net.Dial("tcp", "localhost:8081")

		if err != nil {
			fmt.Println("error listening")
			return
		}

		fmt.Println("connected")

		go clientHandleConnection(conn)
		for {
			if _, err := conn.Write([]byte("hello world")); err != nil {
				fmt.Println("error writing")
				return
			}

			time.Sleep(50000)
		}
	} else {
		ln, err := net.Listen("tcp", ":8081")

		if err != nil {
			fmt.Println("error starting")
			return
		}

		for {
			conn, err := ln.Accept()

			if err != nil {
				log.Fatal("error accepting")
				return
			}

			fmt.Println("someone connect")

			go serverHandleConnection(conn)
		}
	}

}

func clientHandleConnection(conn net.Conn) {
	text := ""
	buf := make([]byte, 1)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("client disconnected")
			} else {
				log.Println("error reading:", err)
			}
			return
		}

		text = string(buf)
		fmt.Println(text)
	}
}

func serverHandleConnection(conn net.Conn) {
	text := ""
	buf := make([]byte, 1)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("client disconnected")
			} else {
				log.Println("error reading:", err)
			}
			return
		}

		text = string(buf)
		fmt.Printf(text)
	}
}
