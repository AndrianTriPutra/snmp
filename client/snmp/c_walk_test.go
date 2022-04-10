package snmp

import (
	"log"
	"testing"
	"time"
)

func Test_Walk(t *testing.T) {
	ip := "127.0.0.1"
	port := uint16(161)
	community := "public"
	timeout := 300 * time.Millisecond
	oid := ".1.3.6.1.2.1.1"

	ans := SnmpWalk(ip, port, community, timeout, oid)
	if ans != nil {
		t.Errorf("error snmpwalk :%v", ans)
	} else {
		log.Println("succes snmpwalk")
	}
}
