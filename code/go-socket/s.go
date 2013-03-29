package main

import (
	"fmt"
	"net"
	"os"
)

type User struct {
	Addr       net.Addr
	Conn       net.Conn
	Postbox    chan string
	QuitSignal chan bool
}

var users []*User

func start(service string) {
	tcpAddr, err := net.ResolveTCPAddr("ipv4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("server start ...")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		user := &User{
			Addr:    conn.RemoteAddr(),
			Conn:    conn,
			Postbox: make(chan string),
		}
		users = append(users, user)

		go handleClient(user)
	}
}

func handleClient(user *User) {
	for {
		select {
		case p := <-user.Postbox:
			user.Conn.Write([]byte(p + "\n"))
		case q := <-user.QuitSignal:
			if q == true {
				user.Conn.Close()
			}
		}
	}
}

func (user *User) SendMessage(text string) {
	user.Postbox <- text
}

func Broadcast(users []*User, text string) {
	for _, user := range users {
		if user != nil {
			user.SendMessage(text)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		// usage()
		os.Exit(1)
	}
	service := ":" + os.Args[1]

	users = make([]*User, 10)

	go start(service)

	var str string

	for {
		fmt.Scanln(&str)
		Broadcast(users, str)
	}
}
