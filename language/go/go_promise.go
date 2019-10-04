package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Promise struct {
	wg sync.WaitGroup
}

func NewPromise(f func()) *Promise {
	p := &Promise{}
	p.wg.Add(1)
	go func() {
		f()
		p.wg.Done()
	}()
	return p
}

func (p *Promise) Then(f func()) *Promise {
	go func() {
		p.wg.Wait()
		p.wg.Add(1)
		f()
		p.wg.Done()
	}()
	return p
}

func main() {
	c := make(chan int)
	defer close(c)
	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	NewPromise(func() {
		select {
		case c <- 1:
			fmt.Println("chan c add 1")
		}
	}).Then(func() {
		select {
		case c <- 2:
			fmt.Println("chan c add 2")
		}
	})
	for v := range c {
		fmt.Println(v)
	}
	<-sigs
}
