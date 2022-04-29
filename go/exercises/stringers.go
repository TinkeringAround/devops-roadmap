// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var addressValues []string
	for _, v := range ip {
		addressValues = append(addressValues, strconv.Itoa(int(v)))
	}
	return strings.Join(addressValues, ".")
}

func Stringers() {
	fmt.Println("## Stringer Interface Exercise ##")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())
	}
}
