package ipinfo

import (
	"context"
	"fmt"
	"net/netip"
)

// PublicIPv4 gets public IPv4 addr by fetching IP from an external service (ipify).
func PublicIPv4(ctx context.Context) (netip.Addr, error) {
	addr, err := getIpifyIP(ctx, ipifyApiV4)
	return addr, errIpify(err, ipifyApiV4)
}

// PublicIPv6 gets public IPv6 addr, first by testing available non-loopback ipv6 interface, then by fetching IP from an external service (ipify).
// This is a faster (local check for fast fail) and more reliable (actual request over the internet) way of detecting ipv6
func PublicIPv6(ctx context.Context) (netip.Addr, error) {
	err := testInterfaceIPv6()
	if err != nil {
		return netip.Addr{}, fmt.Errorf("test ipv6 interface: %w", err)
	}
	addr, err := getIpifyIP(ctx, ipifyApiV6)
	if err != nil {
		return addr, errIpify(err, ipifyApiV6)
	}
	return addr, nil
}
