package main

import (
	"io"
	"log"
	"net"
	"time"
)

func Clock1(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		//handleConn(conn) // handle one connection at a time. Non concurrent. 
		//NOTE: when multiple clients are connected, one client connection must resolve
		// before the next can be provided info.

		//clock2. Uses goroutines and makes it concurrent. Allows multiple clients to receive the time ticker at once.
		go handleConn(conn)
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for{
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil{
			return // e.g. client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
