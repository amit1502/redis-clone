package server

import (
	"github.com/amit1502/redis-clone/config"
	"io"
	"log"
	"net"
	"strconv"
)

func readCommand(c net.Conn) (string, error) {
	// Todo: Max read in one shot is 512 bytes
	// to allow input > 512 bytes, then repated read until
	// we got EOF or designated delimiter
	buf := make([]byte, 512)
	n, err := c.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer() {
	log.Println("Starting a synchronous TCP server on", config.Host, config.Port)

	// keep track of concurrent users
	var con_clients int

	// listening to the configured host:port
	// create a listener object, and start listening to incoming connections
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	if err != nil {
		panic(err)
	}

	for {

		// blocking call: waiting for the new client to connect
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		con_clients += 1
		log.Println("client connected with address:", c.RemoteAddr(), "concurrent clients", con_clients)

		for {

			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				}
				log.Println("err", err)
			}

			log.Println("command", cmd)
			if err = respond(cmd, c); err != nil {
				log.Print("err write:", err)
			}

		}

	}
}
