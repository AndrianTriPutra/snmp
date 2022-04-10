# example snmp


## install
- i recomeded for you install snmp, you can follow this link
- https://lms.onnocenter.or.id/wiki/index.php/SNMP:_UBUNTU_20.04_instalasi_SNMP
- https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-an-snmp-daemon-and-client-on-ubuntu-18-04
- you can check command line snmp with
- snmpwalk -v2c -c public localhost:161 1.3.6.1.2.1.1
- snmpget -v2c -c public localhost:161 1.3.6.1.2.1.1.1.0

## instruction
- git clone https://github.com/AndrianTriPutra/snmp.git snmp
- this code have two use case, and then please open two terminal

### use case 1
- open terminal a, login as root "sudo su"
- cd snmp server
- go run . listen
- open terminal b, user don't need acces root for client
- cd snmp client
- use command line 
- snmptrap -v2c -c public localhost:166 '' . .1.3.6.1.2.1.1.1.0 s "Hello, this is an SNMP trap."
- check terminal a
- and switch to terminal b, go run . send
- and recheck terminal a

### use case 2
- open terminal a, login as root "sudo su"
- cd snmp server
- go run . agent
- open terminal b, user don't need acces root for client
- cd snmp client
- snmpwalk -v2c -c public localhost:166 .1.3.6.1.2.1.1.1
- go run . get or go run . walk
- and recheck terminal a


## dependency
- github.com/twsnmp/gosnmp
