package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:65533")
	if err != nil {
		fmt.Printf("dial failed, err :%v\n", err)
		return
	}
	// 2、利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin) // 从标准输入中读取
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s) //去掉空格
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 3、给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		// 4、从服务端接受回复
		var buf [10024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("recv failed, err:%v\n", err)
			return
		}
		fmt.Println("收到服务端回复", string(buf[:n]))
	}
}
