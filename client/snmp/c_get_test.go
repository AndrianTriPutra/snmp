package snmp

import (
	"log"
	"testing"
	"time"
)

func Test_GET(t *testing.T) {
	ip := "127.0.0.1"
	port := uint16(161)
	community := "public"
	timeout := 300 * time.Millisecond
	oids := []string{".1.3.6.1.2.1.1.1.0", ".1.3.6.1.2.1.1.3.0", ".1.3.6.1.2.1.1.5.0", ".1.3.6.1.2.1.1.7.0"}

	ans := SnmpGet(ip, port, community, timeout, oids)
	if ans != nil {
		t.Errorf("error snmpget :%v", ans)
	} else {
		log.Println("succes snmpget")
	}
}
