package snmp

import (
	"log"
	"time"

	"github.com/twsnmp/gosnmp"
)

func SnmpGet(ip string, port uint16, community string, timeout time.Duration, oids []string) error {

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

	result, err := params.Get(oids)
	if err != nil {
		log.Fatalf("Get err: %v", err)
		return err
	}
	getValue(result)

	return nil
}
