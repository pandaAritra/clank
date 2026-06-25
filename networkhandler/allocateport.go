package networkhandler

import (
	"fmt"
	"net"
)

func AllocatePort() *net.UDPConn {
	// 1. Pass ":0" to automatically request a random available port from the OS
	localAddr, _ := net.ResolveUDPAddr("udp", ":0")
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("Error binding to an automatic port: %v\n", err)
		conn.Close()
	}
	return conn
}
