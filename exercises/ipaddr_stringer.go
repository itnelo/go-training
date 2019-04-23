package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (a IPAddr) String() string {
	var ipAddrStrParts []string

	for _, b := range a {
		numberStr := strconv.Itoa(int(b))

		ipAddrStrParts = append(ipAddrStrParts, numberStr)
	}

	return strings.Join(ipAddrStrParts, ".")
}

func ipaddrStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
