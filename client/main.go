package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	CLIENT_IP = "192.168.1.11"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
)

func sendTg(tg TamGiac, conn net.Conn) {
	// Encode data to JSON
	jsonData, err := json.Marshal(tg)
	if err != nil {
		fmt.Println("Error encoding JSON from client:", err)
		return
	}

	// Send data to the server
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending data to server:", err)
		return
	}
}

func recvTg(conn net.Conn) {

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Client closed the connection.")
			return
		}
		fmt.Println("Error reading data from the client:", err)
		return
	}

	// Unmarshal JSON data received from client
	var tg_kq TamGiac_kq
	if err := json.Unmarshal(buffer[:n], &tg_kq); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the received JSON data from server
	fmt.Printf("Received JSON data from server: %+v\n", tg_kq)

}

func main() {

	// Connect to the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Read and process data from the server
	tg := TamGiac{
		A: 8,
		B: 7,
		C: 9,
	}

	sendTg(tg, conn)

	recvTg(conn)

}
