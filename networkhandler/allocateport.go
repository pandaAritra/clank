package networkhandler

import (
	"fmt"
	"net"
)

func AllocatePort() *net.UDPConn {
	localAddr, _ := net.ResolveUDPAddr("udp", ":0") // ":0" to request  port
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("Error binding to an automatic port: %v\n", err)
		conn.Close()
		panic("connection failed")
	}
	_, allocatedPort, _ := net.SplitHostPort(conn.LocalAddr().String())
	fmt.Printf("[System] OS assigned local port: %s\n", allocatedPort)

	return conn
}
