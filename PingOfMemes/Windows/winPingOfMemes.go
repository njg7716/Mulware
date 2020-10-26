package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/net/icmp"
)

func main() {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := make([]byte, 44)
		_, _, err := conn.ReadFrom(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//Do Things!
		cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FZNIKKF7Q5wA%2Fmaxresdefault.jpg&f=1&nofb=1")
		cmd.Run()
		//time.Sleep(3 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
	}
}
