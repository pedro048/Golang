package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func f1(quit chan bool) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-quit:
				sig := <-sigs
				fmt.Println()
				fmt.Println(sig)
				done <- true
				break
			}
		}
	}()
}

func main() {
	quit := make(chan bool)
	fmt.Println("awaiting signal")
	f1(quit)
	fmt.Println("exiting")
}
