package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (a IPAddr) String() string {
	var strs []string

	for _, b := range a {
		strs = append(strs, strconv.Itoa(int(b)))
	}

	return strings.Join(strs, ".")
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
