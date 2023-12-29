package main

import (
	"fmt"
	"net"
)

// udp server

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 65533,
	})
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read from udp failed, err: %v\n", err)
		}
		fmt.Println("接收到的数据", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to udp failed, err: %v\n", err)
			return
		}
	}
}
