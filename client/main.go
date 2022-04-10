package main

import (
	"client/snmp"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("      go run . send \n")
		fmt.Printf("      go run . get \n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "send":
		sendtrap()

	case "get":
		get()
	}

}

func sendtrap() {
	ip := "127.0.0.1"
	port := uint16(166)
	community := "public"
	timeout := 300 * time.Millisecond
	oid := ".1.3.6.1.2.1.1.1.0"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//ts := time.Now().UTC()
	ts := time.Now().In(loc)
	value := ts.Format(time.RFC3339)

	err := snmp.SendTrap(ip, port, community, timeout, oid, value)
	if err != nil {
		log.Fatalf("error snmpsendtrap :%v", err)
	} else {
		log.Println("succes snmpsendtrap")
	}
}

func get() {
	ip := "127.0.0.1"
	port := uint16(166)
	community := "public"
	timeout := 300 * time.Millisecond
	oids := []string{".1.3.6.1.2.1.1.1.0"}

	err := snmp.SnmpGet(ip, port, community, timeout, oids)
	if err != nil {
		log.Fatalf("error snmpget :%v", err)
	} else {
		log.Println("succes snmpget")
	}
}
