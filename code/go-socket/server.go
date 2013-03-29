package main 

import (
  "fmt" 
  "os"
  "net"
)

type Room struct {
	Name string 
	Users []*User
	Broadcast chan *Message 
}

type User struct {
	Room *Room
	// Name string 
	Connection net.Conn 
	Send chan *Message
}

type Message struct {
	Content string 
}

var ActiveRoom *Room 

func init() {
	ActiveRoom = &Room{
		Name: "ruby",
		Users: make([]*User, 100) ,
		Broadcast: make(chan *Message),
	}

	// fmt.Println(ActiveRoom.Name)

	go ActiveRoom.run()
}


func start(service string){
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

		fmt.Println(conn.RemoteAddr())

		user := &User{Connection: conn, Room: nil}
		ActiveRoom.add(user)
		fmt.Println(ActiveRoom.Users)

		go user.handleClient()
	}
}


func (this *Room) run() {
	for {
		select {
			case b := <- this.Broadcast: 
				for _, user := range this.Users {
					user.Send <- b  
				}
		}
	}
}

func (this *Room) add(user *User) {
	user.Room = this 
	this.Users = append(this.Users, user)
}

func (this *User) handleClient() {
	// defer this.Connection.Close()
	var buf []byte
  this.Connection.Write([]byte(this.Room.Name + "\n"))
  // this.Connection.Write([]byte("hello world\n"))

	for {
  // 有信息时, 读取发送广播
		_, err := this.Connection.Read(buf[0:])
		if err != nil {
			return
		}

		message := &Message{Content: string(buf)}
		this.Room.Broadcast <- message
  
    // 收到广播时, 发到客户端
		select {
			case m :=  <- this.Send:
				this.Connection.Write([]byte(m.Content))
		}
	}
}

func usage() {
	fmt.Printf("usage: %s port \n", os.Args[0] )
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	service := ":" + os.Args[1]

	start(service)
}
