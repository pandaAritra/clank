package networkhandler

import (
	"fmt"
	"net"
	"time"

	"github.com/pion/stun"
)

func Findpublicip(conn *net.UDPConn) {
	stunAddr, _ := net.ResolveUDPAddr("udp", "stun.l.google.com:19302")
	fmt.Println("Connecting to Google STUN to discover public mapping...")

	publicAddr := getPublicIPViaListen(conn, stunAddr)

	printAdder(publicAddr)

}

func getPublicIPViaListen(conn *net.UDPConn, stunAddr *net.UDPAddr) string {
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	_, err := conn.WriteToUDP(message.Raw, stunAddr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic("STUN server request timed out or failed")
		}

		if addr.IP.Equal(stunAddr.IP) || addr.Port == stunAddr.Port {
			res := &stun.Message{Raw: buf[:n]}
			if err := res.Decode(); err != nil {
				panic(err)
			}

			var xorAddr stun.XORMappedAddress
			if err := xorAddr.GetFrom(res); err == nil {
				_ = conn.SetReadDeadline(time.Time{})
				return xorAddr.String()
			}
		}
	}
}

func printAdder(publicAddr string) {

	fmt.Println("\n==================================================")
	fmt.Printf(" YOUR PUBLIC ENDPOINT: %s\n", publicAddr)
	fmt.Println("==================================================")
	fmt.Println("1. Share the endpoint above with C2.")
	fmt.Println("2. Get C2's public endpoint.")
	fmt.Print("\nPaste C2's public endpoint here (IP:port): ")
}
