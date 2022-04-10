package snmp

import (
	"log"
	"os"
	"time"

	"github.com/twsnmp/gosnmp"
)

var OID = [4]string{".1.3.6.1.2.1.1.1.0",
	".1.3.6.1.2.1.1.1.1",
	".1.3.6.1.2.1.1.1.2",
	".1.3.6.1.2.1.1.1.3"}

var Data struct {
	Zero  string
	One   int
	Two   uint32
	Three float32
}

func Agent(ip string, port uint16, community string, timeout time.Duration) {
	Data.One = -1000
	Data.Two = 1000
	Data.Three = 3.14

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
	a.AddMibList(OID[0], gosnmp.OctetString, getZero)
	a.AddMibList(OID[1], gosnmp.Integer, getOne)
	a.AddMibList(OID[2], gosnmp.Counter32, getTwo)
	a.AddMibList(OID[3], gosnmp.OpaqueFloat, getThree)
}

func getZero(oid string) interface{} {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//ts := time.Now().UTC()
	ts := time.Now().In(loc)
	Data.Zero = ts.Format(time.RFC3339)

	log.Printf("[OID:%s] [value:%s]", OID[0], Data.Zero)
	return Data.Zero
}

func getOne(oid string) interface{} {
	Data.One += 1

	log.Printf("[OID:%s] [value:%v]", OID[1], Data.One)
	return Data.One
}

func getTwo(oid string) interface{} {
	Data.Two += 1

	log.Printf("[OID:%s] [value:%v]", OID[1], Data.Two)
	return Data.Two
}

func getThree(oid string) interface{} {
	Data.Three += 0.1

	log.Printf("[OID:%s] [value:%f]", OID[1], Data.Three)
	return Data.Three
}
