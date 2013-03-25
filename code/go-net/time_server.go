package main 

import (
	"fmt"
	"time"
	"net"
	"os"
)	

func main() {
	// 要监听的地址,
	service := ":1200"
	// 返回一个tcpAddr 结构
	tcpAddr, err := net.ResolveTCPAddr("ipv4", service)
	checkError(err)

	// 建立监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// 有连接来
		conn, err := listener.Accept()
		if err != nil {
			continue 
		}
		fmt.Println(conn.RemoteAddr())

		daytime := time.Now().String()
		// 写入数据
		conn.Write([]byte(daytime))
		conn.Close()
	}


}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
                os.Exit(1)
        }
}