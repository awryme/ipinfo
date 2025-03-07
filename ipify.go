package ipinfo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"time"
)

const ipifyApiV4 = "https://api.ipify.org"
const ipifyApiV6 = "https://api6.ipify.org"

var HttpTimeout = time.Second * 30

func errIpify(err error, addr string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("get ip from %s: %w", addr, err)
}

func getIpifyIP(ctx context.Context, apiAddr string) (addr netip.Addr, err error) {
	cl := &http.Client{
		Timeout: HttpTimeout,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiAddr, nil)
	if err != nil {
		return addr, fmt.Errorf("create http request: %w", err)
	}
	res, err := cl.Do(req)
	if err != nil {
		return addr, fmt.Errorf("do http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return addr, fmt.Errorf("bad status code %d", res.StatusCode)
	}

	ipBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return addr, fmt.Errorf("read request body: %w", err)
	}
	return netip.ParseAddr(string(ipBytes))
}
