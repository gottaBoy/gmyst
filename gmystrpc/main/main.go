package main

import (
	"encoding/json"
	"fmt"
	"gmyst/gmystrpc"
	"gmyst/gmystrpc/codec"
	"log"
	"net"
	"time"
)

//在 startServer 中使用了信道 addr，确保服务端端口监听成功，客户端再发起请求。
//客户端首先发送 Option 进行协议交换，接下来发送消息头 h := &codec.Header{}，和消息体 gmystrpc req ${h.Seq}。
//最后解析服务端的响应 reply，并打印出来
func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	gmystrpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple gmystrpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(gmystrpc.DefaultOption)
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("gmystrpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}