package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SleepAndWrite() {
	min := 5
	max := 10
	rand.Seed(time.Now().UTC().UnixNano())
	sleepTime := rand.Intn(max-min) + min
	fmt.Println("Slept for ", sleepTime, " seconds")
	time.Sleep(sleepTime * time.Second)
}

func main() {
	SleepAndWrite()
}
