package net

import (
	"fmt"
	"net"
)

//tcpReadData
/*
	step
		1: bind
		2: listen
		3: accept //accept
*/
const Port = 8989

func tcpServerDealData1() {

	handleC := func(c net.Conn) error { //如果我始终没有处理 这个connection 会发生什么问题
		//read data
		//logic deal
		//.......
		//write data
		return nil
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleC(conn)
	}
}

func tcpClientSendData1() {
	//conn, err := net.Dial("tcp", fmt.Sprintf(":%d", Port))
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
}
