package main

import (
	"flag"
	"fmt"
	"os"
	"server/snmp"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("      go run . listen \n")
		fmt.Printf("      go run . agent \n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "listen":
		listen()

	case "agent":
		agent()
	}

}

func listen() {
	ip := "127.0.0.1"
	port := uint16(166)
	community := "public"
	timeout := 300 * time.Millisecond
	snmp.Listen("127.0.0.1:166", ip, port, community, timeout)
}

func agent() {
	ip := "127.0.0.1"
	port := uint16(166)
	community := "public"
	timeout := 300 * time.Millisecond
	snmp.Agent(ip, port, community, timeout)
}
