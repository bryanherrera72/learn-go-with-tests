package main

// just use this to read tcp info if you don't have a tool like nc or telnet
import (
	"io"
	"log"
	"net"
	"os"
)

func Netcat1(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}