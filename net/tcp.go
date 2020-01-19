package net

import (
	"fmt"
	"net"
)

func ListenSetup() {
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err: ", err)
			break
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	packet := make([]byte, 1024)
	for {
		_, _ = conn.Read(packet)
		_, _ = conn.Write(packet)
	}
}
