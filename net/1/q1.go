package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8888"
	SERVER_TYPE = "tcp"
)

func main() {
	server := func() {
		// Server
		fmt.Println("Server Running...")
		server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		defer server.Close()
		fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
		fmt.Println("Waiting for client...")
		for {
			connection, err := server.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			fmt.Println("client connected")
			go processClient(connection)
		}
	}
	client1 := func() {
		// Client 1
		//establish connection
		rand.Seed(time.Now().UnixNano())
		connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
		if err != nil {
			panic(err)
		}
		///send some data
		arrayMsg1 := [2]string{"Hello", "Hi"}
		randNumber1 := rand.Intn(1-0+1) + 0
		_, err = connection.Write([]byte(arrayMsg1[randNumber1]))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received for Client 1: ", string(buffer[:mLen]))
		rand1_sleep := rand.Intn(6) + 2
		time.Sleep(time.Duration(rand1_sleep) * time.Second)
		defer connection.Close()
	}
	client2 := func() {
		// Client 2
		//establish connection
		rand.Seed(time.Now().UnixNano())
		connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
		if err != nil {
			panic(err)
		}
		///send some data
		arrayMsg2 := [2]string{"Hello", "Hi"}
		randNumber2 := rand.Intn(1-0+1) + 0
		_, err = connection.Write([]byte(arrayMsg2[randNumber2]))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received for Client 2: ", string(buffer[:mLen]))
		rand2_sleep := rand.Intn(6) + 2
		time.Sleep(time.Duration(rand2_sleep) * time.Second)
		defer connection.Close()
	}

	go server()
	go client1()
	go client2()
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received for Server: ", string(buffer[:mLen]))
	if string(buffer[:mLen]) == "Hello" {
		_, err = connection.Write([]byte("Hello World"))
	} else {
		_, err = connection.Write([]byte("I didn't understand"))
	}

	connection.Close()
}
