// chatroom server
package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConns = make(map[string]net.Conn)//存储客户端链接映射 key为链接ip:port value为连接对象conn
var messageQueue = make(chan string, 1000)//消息队列 带缓冲的buf

var quitChan = make(chan bool)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//消息接收协程
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		if numOfBytes != 0 {
			/*结尾buf[0:numOfBytes]的原因是：numOfBytes是指接收的字节数 如果只用string(buf)
	    		可能会导致接收字符串中有其他之前接收的字符
	    	*/
			message := string(buf[0:numOfBytes])

			//将消息放入到消息队列
			messageQueue <- message
		}
	}
}

//消费者协程
func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

//消息解析函数
func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		sendMessage := contents[1]

		addr = strings.Trim(addr, " ")

		//通过addr查看是否有链接对象
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
			}
		}
	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()

	fmt.Println("Server is waiting....")

	go ConsumeMessage()

	for {
		conn, err := listen_socket.Accept()
		CheckError(err)

		//将conn存储到onlineConns映射表中
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for i := range onlineConns {
			fmt.Println(i)
		}

		go ProcessInfo(conn)
	}
}
