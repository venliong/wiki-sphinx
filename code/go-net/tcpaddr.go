package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// tcpaddr  www.google.com:80 
		fmt.Fprintf(os.Stderr, "usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// 函数ResolveTCPAddr用来创建一个TCPAddr

	// type TCPAddr struct {
	//		 IP   IP
	// 		 Port int
	// }
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// ni 表示本地地址, 所以设置成nil, tcpAddr 是一个远程地址, tcp表示使用的协议
	// 建立连接 
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	//写入数据 
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 在这种情况下，读本质上是一个单一的来自服务器的响应，这将终止文件结束的连接。
	// 但是，它可能包括多个TCP数据包，所以我们需要不断地读，直到文件的末尾。
	// 在io/ioutil下的ReadAll函数考虑这些问题，并返回完整响应。

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
