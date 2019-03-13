// chatroom server
package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//专门负责接收信息的协程
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		//如果接收字节数不为0 说明有消息发过来
		if numOfBytes != 0 {
			fmt.Printf("Has received this message: %s\n", string(buf))
		}
	}
}

func main() {
	//开启监听socket 监听在本地8080端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()

	fmt.Println("Server is waiting....")
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)

		//如果有客户端链接 则打开一个协程处理
		go ProcessInfo(conn)
	}
}
