package utils

import (
	"context"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

func UseProxyClient() (*http.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", os.Getenv("PROXY"), nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// add more handling here
			return dialer.Dial(network, addr)
		},
	}
	client := &http.Client{Transport: transport}

	return client, nil
}
