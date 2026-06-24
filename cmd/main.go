package main

import (
	"fmt"
	"net"
)

func main() {
	network, adress := net.ResolveUDPAddr("udp", ":0")
	// conn, err := net.ListenUDP("udp", network)
	fmt.Println(network, adress)
	// 1. Bind a completely un-connected, generic local UDP listener
	// 	localAddr, _ := net.ResolveUDPAddr("udp", localPort)
	// 	conn, err := net.ListenUDP("udp", localAddr)
	// 	if err != nil {
	// 		fmt.Printf("Error binding to port: %v\n", err)
	// 		return
	// 	}
	// 	defer conn.Close()

	// 	// 2. Query Google STUN by passing the packet explicitly over the listener
	// 	stunAddr, _ := net.ResolveUDPAddr("udp", "stun.l.google.com:19302")
	// 	fmt.Println("Connecting to Google STUN to discover public mapping...")
	// 	publicAddr := getPublicIPViaListen(conn, stunAddr)

	// 	fmt.Println("\n==================================================")
	// 	fmt.Printf(" YOUR PUBLIC ENDPOINT: %s\n", publicAddr)
	// 	fmt.Println("==================================================")
	// 	fmt.Println("1. Share the endpoint above with C2.")
	// 	fmt.Println("2. Get C2's public endpoint.")
	// 	fmt.Print("\nPaste C2's public endpoint here (IP:port): ")

	// 	var remoteAddrStr string
	// 	fmt.Scanln(&remoteAddrStr)

	// 	remoteAddr, err := net.ResolveUDPAddr("udp", remoteAddrStr)
	// 	if err != nil {
	// 		fmt.Printf("Invalid remote address: %v\n", err)
	// 		return
	// 	}

	// 	fmt.Printf("\n[Starting Hole Punch] Blasting 'alive' to %s every 5s...\n", remoteAddrStr)

	// 	// 3. HEARTBEAT GENERATOR: This will now work cleanly because the socket is not locked!
	// 	go func() {
	// 		ticker := time.NewTicker(5 * time.Second)
	// 		for range ticker.C {
	// 			_, err := conn.WriteToUDP([]byte("alive"), remoteAddr)
	// 			if err != nil {
	// 				fmt.Printf("\n[Error sending]: %v", err)
	// 			}
	// 		}
	// 	}()

	// 	// 4. RECEIVER: Listens for incoming 'alive' heartbeats from your peer
	// 	buf := make([]byte, 1024)
	// 	for {
	// 		n, addr, err := conn.ReadFromUDP(buf)
	// 		if err != nil {
	// 			continue
	// 		}

	// 		payload := string(buf[:n])
	// 		if payload == "alive" {
	// 			currentTime := time.Now().Format("15:04:05")
	// 			fmt.Printf("[%s] -> Hole Punched! Received 'alive' heartbeat from %s\n", currentTime, addr)
	// 		}
	// 	}
	// }

	// // Fixed STUN lookup that works cleanly with net.ListenUDP without pre-connecting
	// func getPublicIPViaListen(conn *net.UDPConn, stunAddr *net.UDPAddr) string {
	// 	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	// 	// Send the binding request directly to Google STUN address
	// 	_, err := conn.WriteToUDP(message.Raw, stunAddr)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Listen specifically for Google's response back
	// 	buf := make([]byte, 1024)
	// 	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // safety timeout

	// 	for {
	// 		n, addr, err := conn.ReadFromUDP(buf)
	// 		if err != nil {
	// 			panic("STUN server request timed out or failed")
	// 		}

	// 		// Make sure the packet came from Google STUN
	// 		if addr.IP.Equal(stunAddr.IP) || addr.Port == stunAddr.Port {
	// 			res := &stun.Message{Raw: buf[:n]}
	// 			if err := res.Decode(); err != nil {
	// 				panic(err)
	// 			}

	//			var xorAddr stun.XORMappedAddress
	//			if err := xorAddr.GetFrom(res); err == nil {
	//				_ = conn.SetReadDeadline(time.Time{}) // Clear timeout for normal use
	//				return xorAddr.String()
	//			}
	//		}
	//	}
}
