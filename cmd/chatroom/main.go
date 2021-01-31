package main

import (
	"chatroom/config"
	"chatroom/server"
	"fmt"
)

func main() {
	fmt.Println("welecome to chatroom:", config.Server.Addr)
	server.Start(config.Server.Addr)
}
