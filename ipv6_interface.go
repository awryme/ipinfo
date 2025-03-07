package ipinfo

import (
	"fmt"
	"net"
	"strings"
)

var ErrNoIPv6Interface = fmt.Errorf("no ipv6 interfaces found")

// HasInterfaceIPv6 checks if network has public (not loopback) ipv6 interface
func testInterfaceIPv6() error {
	ifs, err := net.Interfaces()
	if err != nil {
		return fmt.Errorf("list interfaces: %w", err)
	}

	for _, iface := range ifs {
		// ignore loopbacks
		if iface.Flags&net.FlagLoopback == 1 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return fmt.Errorf("get interface %s addrs: %w", iface.Name, err)
		}
		for _, addr := range addrs {
			// just count colons in a string
			// even with port ipv4 can only have single ':'
			// ipv6 must have at least 2
			if strings.Count(addr.String(), ":") >= 2 {
				return nil
			}
		}
	}

	return ErrNoIPv6Interface
}
