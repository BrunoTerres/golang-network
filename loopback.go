package main

import (
	"github.com/jsimonetti/rtnetlink/rtnl"
	"log"
	"net"
	//"log"
	//"net"
)

func main() {
	var loopback *net.Interface

	conn, err := rtnl.Dial(nil)
	if err != nil {
		return
	}
	defer conn.Close()

	links, err := conn.Links()

	for _, l := range links {
		if l.Name == "lo" {
			loopback = l
			log.Printf("Name1: %s, Flags: %s\n\n", l.Name, l.Flags)
		}
		conn.LinkDown(loopback)
		loopback, _ = conn.LinkByIndex(loopback.Index)
		log.Printf("Name2: %s, Flags: %s\n", loopback)

		conn.LinkUp(loopback)
		loopback, _ = conn.LinkByIndex(loopback.Index)
		log.Printf("Name3: %S, Flags: %s\n", loopback.Name, loopback.Flags)
	}
}
