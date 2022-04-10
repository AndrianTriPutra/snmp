package main

import (
	"client/snmp"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var Setting struct {
	IP        string
	Port      uint16
	Community string
	Timeout   time.Duration
}

func main() {
	Setting.IP = "127.0.0.1"
	Setting.Port = 166
	Setting.Community = "public"
	Setting.Timeout = 300 * time.Millisecond

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("      go run . send \n")
		fmt.Printf("      go run . walk \n")
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

	case "walk":
		walk()

	case "get":
		get()
	}

}

func sendtrap() {
	oid := ".1.3.6.1.2.1.1.1.0"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//ts := time.Now().UTC()
	ts := time.Now().In(loc)
	value := ts.Format(time.RFC3339)

	err := snmp.SendTrap(Setting.IP, Setting.Port, Setting.Community, Setting.Timeout, oid, value)
	if err != nil {
		log.Fatalf("error snmpsendtrap :%v", err)
	} else {
		log.Println("succes snmpsendtrap")
	}
}

func get() {
	//oids := []string{".1.3.6.1.2.1.1.1.0"}
	oids := []string{".1.3.6.1.2.1.1.1.0",
		".1.3.6.1.2.1.1.1.1",
		".1.3.6.1.2.1.1.1.2",
		".1.3.6.1.2.1.1.1.3"}

	err := snmp.SnmpGet(Setting.IP, Setting.Port, Setting.Community, Setting.Timeout, oids)
	if err != nil {
		log.Fatalf("error snmpget :%v", err)
	} else {
		log.Println("succes snmpget")
	}
}

func walk() {
	oid := ".1.3.6.1.2.1.1.1"

	err := snmp.SnmpWalk(Setting.IP, Setting.Port, Setting.Community, Setting.Timeout, oid)
	if err != nil {
		log.Fatalf("error snmpwalk :%v", err)
	} else {
		log.Println("succes snmpwalk")
	}
}
