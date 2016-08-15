package debugutils

import (
	"fmt"
	"net"
	"os"
	"time"
)

const Port = ":8081"

func EnableShutdownUtility() {
	go listenOnPort()
}

func listenOnPort() {
	ln, err := net.Listen("tcp", Port)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}
}

func ShutdownRunningServer() {
	net.Dial("tcp", "127.0.0.1" + Port)
	time.Sleep(50 * time.Millisecond)
}
