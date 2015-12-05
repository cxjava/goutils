package httputils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hailiang/socks"
)

// DefaultClient returns a new http.Client with the same default values as
// http.Client, but with a non-shared Transport
func DefaultClientByProxy(proxyStr string) *http.Client {
	return &http.Client{
		Transport: DefaultTransportByProxy(proxyStr),
	}
}

// DefaultTransportByProxy returns a new http.Transport with the same default values
// as http.DefaultTransport
func DefaultTransportByProxy(proxyStr string) (transport *http.Transport) {
	transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		TLSHandshakeTimeout: 10 * time.Second,
	}

	if len(proxyStr) > 0 {
		proxyURL := strings.ToUpper(proxyStr)
		fmt.Println("proxy URL:", proxyURL)
		if strings.HasPrefix(proxyURL, "SOCKS") {
			//proxyURL like :"SOCKS5://127.0.0.1:1080" or "SOCKS://127.0.0.1:9150"
			switch proxyURL[:6] {
			case "SOCKS5":
				dialSocks5Proxy := socks.DialSocksProxy(socks.SOCKS5, proxyURL[9:])
				transport.Dial = dialSocks5Proxy
			case "SOCKS4":
				dialSocks4Proxy := socks.DialSocksProxy(socks.SOCKS4, proxyURL[9:])
				transport.Dial = dialSocks4Proxy
			default:
				proxyURL = strings.Replace(proxyURL, "SOCKS://", "", -1)
				dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, proxyURL)
				transport.Dial = dialSocksProxy
			}
		} else {
			pu, err := url.Parse(proxyURL)
			if err != nil {
				fmt.Println("DefaultTransportByProxy url.Parse:", err, proxyURL)
				return
			}
			transport.Proxy = http.ProxyURL(pu)
		}
	}
	return
}
