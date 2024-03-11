package main

import (
	"E2EE-Chat/cmd/client"
	"E2EE-Chat/cmd/server"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	server.Server()
	client.Client()
}
