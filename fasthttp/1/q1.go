package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"
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
		server, err := fasthttp.ListenAndServe(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
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
		_, err = connection.Write([]byte("msg"))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received for Client 1: ", string(buffer[:mLen]))
		rand1_sleep := rand.Intn(6) + 2
		time.Sleep(time.Duration(rand1_sleep) * time.Second)

		resp, err := http.Get("http://localhost:8888/")

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))

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
		_, err = connection.Write([]byte("msg"))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received for Client 2: ", string(buffer[:mLen]))
		rand2_sleep := rand.Intn(6) + 2
		time.Sleep(time.Duration(rand2_sleep) * time.Second)

		resp, err := http.Get("http://localhost:8888/")

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))

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
	_, err = connection.Write([]byte("<html><head><title>Title</title></head><body><h1>Test!</h1></body></html>"))

	connection.Close()
}
