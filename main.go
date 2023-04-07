package main

import (
	"ProtogeCell/config"
	"ProtogeCell/server"
)

func main() {
	config.GoogleConfig()
	server.ConnectDB()
	server.Serve()

}
