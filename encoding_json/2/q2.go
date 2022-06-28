package main

import (
	"encoding/json"
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

type User1 struct {
	Name1 string `json:"text1"`
}
type User2 struct {
	Name2 string `json:"text2"`
}
type User3 struct {
	Name3 string `json:"text3"`
}
type User4 struct {
	Name4 string `json:"text4"`
}

type MyResponse struct {
	Name string
}

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
		_, err = connection.Write([]byte("http://localhost:8888/"))
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
		_, err = connection.Write([]byte("http://localhost:8888/test01"))
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

	u1 := User1{}
	u2 := User2{}
	u3 := User3{}
	u4 := User4{}

	b1 := []byte(`{"text1": "Test!"}`)
	b2 := []byte(`{"text2": "Test01!"}`)
	b3 := []byte(`{"text3": "Test02!"}`)
	b4 := []byte(`{"text4": "Default!"}`)

	err1 := json.Unmarshal(b1, &u1)
	err2 := json.Unmarshal(b2, &u2)
	err3 := json.Unmarshal(b3, &u3)
	err4 := json.Unmarshal(b4, &u4)

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)

	if string(buffer[:mLen]) == "http://localhost:8888/" {
		var myResponse1 = MyResponse{Name: u1.Name1}
		_, err = connection.Write([]byte(myResponse1.Name))
	} else if string(buffer[:mLen]) == "http://localhost:8888/test01" {
		var myResponse2 = MyResponse{Name: u2.Name2}
		_, err = connection.Write([]byte(myResponse2.Name))
	} else if string(buffer[:mLen]) == "http://localhost:8888/test02" {
		var myResponse3 = MyResponse{Name: u3.Name3}
		_, err = connection.Write([]byte(myResponse3.Name))
	} else {
		var myResponse4 = MyResponse{Name: u4.Name4}
		_, err = connection.Write([]byte(myResponse4.Name))
	}

	connection.Close()
}
