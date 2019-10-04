package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/lcullen/mardown/framwork/thrift_demo/gen-go/batu/demo"
)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)

func main() {
	for i := 0; i < 10; i++ {
		perform(i)
		time.Sleep(1 * time.Minute)
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func perform(i int) {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := demo.NewBatuThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	paramMap := make(map[string]string)
	paramMap["a"] = "batu.demo"
	paramMap["b"] = "test" + strconv.Itoa(i+1)
	r1, err := client.CallBack(time.Now().Unix(), "go client", paramMap)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("GOClient Call->", r1)

	model := demo.Article{1, "Go第一篇文章", "我在这里", "liuxinming"}
	client.Put(&model)
	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))

}
