package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c := make(chan os.Signal)

	go func() {
		for {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r1_sleep := rand.New(s1)
			rand1 := r1.Intn(100)
			rand1_sleep := r1_sleep.Intn(6) + 2
			time.Sleep(time.Duration(rand1_sleep) * time.Second)
			c1 <- rand1 // receive a value from c1 channel
		}
	}()

	go func() {
		for {
			s2 := rand.NewSource(time.Now().UnixNano())
			r2 := rand.New(s2)
			r2_sleep := rand.New(s2)
			rand2 := r2.Intn(100)
			rand2_sleep := r2_sleep.Intn(6) + 2
			time.Sleep(time.Duration(rand2_sleep) * time.Second)
			c2 <- rand2 // receive a value from c2 channel
		}
	}()

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for i := 0; i < 20; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("The producer 01 sent", msg1)
			break
		case msg2 := <-c2:
			fmt.Println("The producer 02 sent", msg2)
			break
		}
	}
}
