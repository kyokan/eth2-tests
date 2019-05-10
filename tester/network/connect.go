package network

import (
	"log"
	"net"
	"strconv"
	"time"
)

func Connect(host string, port int) error {
	log.Printf("Connecting to %s", host+":"+strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Second)
	if conn != nil {
		log.Printf("Connection established with remote host at %s:%d", host, port)
		defer conn.Close()
	}
	return err
}
