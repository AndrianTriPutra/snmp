package snmp

import (
	"log"
	"net"
	"time"

	"github.com/twsnmp/gosnmp"
)

func Listen(host string, ip string, port uint16, community string, timeout time.Duration) {

	tl := gosnmp.NewTrapListener()
	tl.OnNewTrap = trapHandler
	params := &gosnmp.GoSNMP{
		Target:    ip,
		Port:      port,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   timeout,
		//Logger:    gosnmp.NewLogger(log.New(os.Stdout, "", 0)), //optional
	}

	//tl.Params = gosnmp.Default
	tl.Params = params

	log.Println("starting listen . . .")
	err := tl.Listen(host)
	if err != nil {
		log.Panicf("error in listen: %s", err)
	}
}

func trapHandler(packet *gosnmp.SnmpPacket, addr *net.UDPAddr) {
	for _, v := range packet.Variables {
		switch v.Type {
		case gosnmp.OctetString:
			oid := string(v.Name)
			value := string(v.Value.([]byte))
			log.Printf("[string] [oid:%s] [value:%s]", oid, value)
		}
	}
}
