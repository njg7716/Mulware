package main

import (
	"fmt"
	"log"

	"golang.org/x/net/icmp"
)

func main() {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := make([]byte, 4)
		length, sourceIP, err := conn.ReadFrom(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//Do Things!
		fmt.Printf("message = '%s', length = %d, source-ip = %s\n", string(msg), length, sourceIP)
	}
	_ = conn.Close()
}
