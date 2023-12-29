package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//udp client demo

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 65533,
	})
	if err != nil {
		fmt.Printf("dial failed, err :%v\n", err)
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Printf("send to server failed, err :%v\n", err)
		}
		// 接收数据
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("recv from server failed, err :%v\n", err)
		}
		fmt.Printf("read from %v, msg:%v\n", addr, string(buf[:n]))

	}
}
