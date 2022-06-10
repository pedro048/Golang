package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SleepAndPrint(num int) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	rand := r.Intn(8) + 2
	fmt.Printf("Runner %d slept for %d seconds.", num, rand)
	fmt.Printf("\n")
	time.Sleep(time.Duration(rand) * time.Second)
	wg.Done()
}

func main() {
	wg.Add(1)
	go SleepAndPrint(1)
	go SleepAndPrint(2)
	go SleepAndPrint(3)
	go SleepAndPrint(4)
	go SleepAndPrint(5)
	go SleepAndPrint(6)
	wg.Wait()
}
