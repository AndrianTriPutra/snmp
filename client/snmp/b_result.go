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

		case gosnmp.Integer:
			fmt.Printf("integer: %v\n", gosnmp.ToBigInt(variable.Value))

		case gosnmp.Counter32:
			fmt.Printf("count32: %v\n", variable.Value)

		case gosnmp.OpaqueFloat:
			fmt.Printf("float: %f\n", variable.Value)

		}
	}
}

func getValue(result *gosnmp.SnmpPacket) {
	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))

		case gosnmp.Integer:
			fmt.Printf("integer: %v\n", gosnmp.ToBigInt(variable.Value))

		case gosnmp.Counter32:
			fmt.Printf("count32: %v\n", variable.Value)

		case gosnmp.OpaqueFloat:
			fmt.Printf("float: %f\n", variable.Value)

		}
	}
}
