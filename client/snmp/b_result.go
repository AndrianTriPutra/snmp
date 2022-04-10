package snmp

import (
	"fmt"

	"github.com/twsnmp/gosnmp"
)

func walkValue(results []gosnmp.SnmpPDU) {
	for i, variable := range results {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))

		default:
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}

func getValue(result *gosnmp.SnmpPacket) {
	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))

		default:
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}
