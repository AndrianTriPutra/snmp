package snmp

import (
	"log"
	"os"
	"time"

	"github.com/twsnmp/gosnmp"
)

var OID = [1]string{".1.3.6.1.2.1.1.1.0"}

func Agent(ip string, port uint16, community string, timeout time.Duration) {
	forever := make(chan os.Signal, 1)

	g := &gosnmp.GoSNMP{}
	g.Port = port
	g.Community = community
	g.Version = gosnmp.Version2c
	g.Timeout = timeout
	a := &gosnmp.GoSNMPAgent{
		Port:   166,
		IPAddr: ip,
		//Logger: log.New(os.Stdout, "", 0),
		Snmp: g,
	}

	handler(a)

	log.Printf("starting agent . . .")
	if err := a.Start(); err != nil {
		log.Fatalf("error agent :%v", err)
	}

	<-forever
}

func handler(a *gosnmp.GoSNMPAgent) {
	a.AddMibList(OID[0], gosnmp.OctetString, getOne)
}

func getOne(oid string) interface{} {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//ts := time.Now().UTC()
	ts := time.Now().In(loc)
	timestamp := ts.Format(time.RFC3339)

	log.Printf("[OID:%s] [value:%s]", OID[0], timestamp)
	return timestamp
}
