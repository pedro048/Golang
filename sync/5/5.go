package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 0
	var m sync.Mutex

	go func() {
		for {
			m.Lock()
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r1_sleep := rand.New(s1)
			rand1 := r1.Intn(100)
			fmt.Println("The producer 01 sent", rand1)
			rand1_sleep := r1_sleep.Intn(8) + 2
			time.Sleep(time.Duration(rand1_sleep) * time.Second)
			n = rand1
			m.Unlock()
		}
	}()

	go func() {
		for {
			m.Lock()
			s2 := rand.NewSource(time.Now().UnixNano())
			r2 := rand.New(s2)
			r2_sleep := rand.New(s2)
			rand2 := r2.Intn(100)
			fmt.Println("The producer 02 sent", rand2)
			rand2_sleep := r2_sleep.Intn(8) + 2
			time.Sleep(time.Duration(rand2_sleep) * time.Second)
			n = rand2
			m.Unlock()
		}
	}()

	go func() {
		for {
			m.Lock()
			consumer1 := n // received for the receiver 01
			fmt.Println("The receiver 01 received", consumer1)
			s4 := rand.NewSource(time.Now().UnixNano())
			r4_sleep := rand.New(s4)
			rand4_sleep := r4_sleep.Intn(8) + 2
			time.Sleep(time.Duration(rand4_sleep) * time.Second)
			m.Unlock()
		}
	}()

	go func() {
		for {
			m.Lock()
			consumer2 := n // received for the receiver 02
			fmt.Println("The receiver 02 received", consumer2)
			s3 := rand.NewSource(time.Now().UnixNano())
			r3_sleep := rand.New(s3)
			rand3_sleep := r3_sleep.Intn(8) + 2
			time.Sleep(time.Duration(rand3_sleep) * time.Second)
			m.Unlock()
		}
	}()
}
