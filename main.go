package main

import (
	"flag"
	"github.com/amit1502/redis-clone/config"
	"github.com/amit1502/redis-clone/server"
	"log"
)

// Task: Function to read command-line flags
func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis server")
	flag.IntVar(&config.Port, "port", 7379, "port for the redis server")
	flag.Parse()
}

func main() {
	setupFlags()

	// read coomand-line flags
	log.Println("starting redis server ...")

	// start the redis server
	server.RunSyncTCPServer()
}
