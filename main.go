package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	var port int
	var ip string
	flag.IntVar(&port,"port",30000,"监听端口,默认值30000")
	flag.StringVar(&ip,"ip","","监听地址")
	flag.Parse()
	conn, err := net.Dial("udp", ip+":"+strconv.Itoa(port))
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("本地ip+port:",conn.LocalAddr())
	fmt.Println("network:",conn.LocalAddr().Network())
	//客户端发起交谈
	_,err=conn.Write([]byte("客户端发送请求"))
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	udpAddr,err:=net.ResolveUDPAddr(conn.LocalAddr().Network(),":"+strings.Split(conn.LocalAddr().String(),":")[1])
	if err!=nil{
		log.Fatal(err)
	}
	serverConn,err:=net.ListenUDP(conn.LocalAddr().Network(),udpAddr)
	if err!=nil{
		log.Fatal(err)
	}
	defer serverConn.Close()

	b:=make([]byte,1024)
	m,err:=conn.Read(b)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(b[:m]))

	//接收服务端消息
	buffer := make([]byte, 1024)
	n, remoteAddress,_ := serverConn.ReadFromUDP(buffer)
	fmt.Println("remoteAddress",remoteAddress)
	fmt.Println("服务端返回:"+string(buffer[:n]))


}