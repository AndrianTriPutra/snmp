package snmp

import (
	"log"
	"time"

	"github.com/twsnmp/gosnmp"
)

func SendTrap(ip string, port uint16, community string, timeout time.Duration, oid, value string) error {
	params := &gosnmp.GoSNMP{
		Target:    ip,
		Port:      port,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   timeout,
		//Logger:    gosnmp.NewLogger(log.New(os.Stdout, "", 0)), //optional
	}
	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect err: %v", err)
		return err
	}
	defer params.Conn.Close()

	pdu := gosnmp.SnmpPDU{
		Name:  oid,
		Type:  gosnmp.OctetString,
		Value: value,
	}

	trap := gosnmp.SnmpTrap{
		Variables: []gosnmp.SnmpPDU{pdu},
	}

	log.Printf("trap :%v", trap)

	_, err = params.SendTrap(trap)
	if err != nil {
		log.Fatalf("SendTrap err: %v", err)
		return err
	}

	return nil
}
