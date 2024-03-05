package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
)

func sendTgKq(tg_kq TamGiac_kq, conn net.Conn) {

	// Encode data to JSON
	jsonDataResult, err := json.Marshal(tg_kq)
	if err != nil {
		fmt.Println("Error encoding JSON from server:", err)
		return
	}
	// Send data result to client
	_, err = conn.Write(jsonDataResult)
	if err != nil {
		fmt.Println("Error sending data to client:", err)
		return
	}

}

func recvTg(conn net.Conn) TamGiac {

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	// Read data from the client
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Client closed the connection.")
		}
		fmt.Println("Error reading data from the client:", err)
	}

	// Unmarshal JSON data received from client
	var tg TamGiac
	if err := json.Unmarshal(buffer[:n], &tg); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return tg
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	tg := recvTg(conn)

	// Print the received JSON data
	fmt.Printf("Received JSON data from client: %+v\n", tg)

	// handle tg_kq
	var tg_kq TamGiac_kq
	tg_kq.Cv = tg.TinhChuVi()
	tg_kq.Dt = tg.TinhDienTich()

	sendTgKq(tg_kq, conn)

}

func main() {

	// Listen for incoming connections
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accept incoming connections:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}
