package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var a []string
	go func() {
		for {
			fmt.Println(11)
			a = append(a, "kkxx")
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}
