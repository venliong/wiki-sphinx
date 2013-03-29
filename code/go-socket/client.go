package main 

import (
	"fmt"
	"net"
	"os"
	// "io/ioutil"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

  var str string 
  
	for {
		fmt.Scanf("%s", &str)
		_, err = conn.Write([]byte(str))
		checkError(err)
	}

	// result, err := ioutil.ReadAll(conn)
	// checkError(err)

	// fmt.Println(string(result))


}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}