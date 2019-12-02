package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func Test_clientFirstShakeHand(t *testing.T) {
	tr := time.AfterFunc(1*time.Hour, func() {

	})
	tr.Stop()
}

func clientFirstShakeHand() {
	//client 可以使用nc 来实现
}

func serverListenAndServe() {
	ls, err := net.Listen("tcp", ":8989")
	if err != nil {
		panic(err)
	}
	defer ls.Close()
	fmt.Println(ls.Addr().String())
	time.Sleep(10 * time.Minute)
}

func unixSocketServ(path string) {
	//net.ListenUnix()
}
