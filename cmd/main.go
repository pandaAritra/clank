package main

import (
	"fmt"
	"net"
	"time"

	"github.com/pandaAritra/clank/networkhandler"
)

func main() {

	conn := networkhandler.AllocatePort()
	defer conn.Close()

	networkhandler.Findpublicip(conn)

	var remoteAddrStr string
	fmt.Scanln(&remoteAddrStr)

	remoteAddr, err := net.ResolveUDPAddr("udp", remoteAddrStr)
	if err != nil {
		fmt.Printf("Invalid remote address: %v\n", err)
		return
	}

	fmt.Printf("\n[Starting Hole Punch] Blasting 'alive' to %s every 5s...\n", remoteAddrStr)

	// 3. HEARTBEAT GENERATOR
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			_, err := conn.WriteToUDP([]byte("alive"), remoteAddr)
			if err != nil {
				fmt.Printf("\n[Error sending]: %v", err)
			}
		}
	}()

	// 4. RECEIVER
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		payload := string(buf[:n])
		if payload == "alive" {
			currentTime := time.Now().Format("15:04:05")
			fmt.Printf("[%s] -> Hole Punched! Received 'alive' heartbeat from %s\n", currentTime, addr)
		}
	}
}
