package createtransport

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
)

func CreateTransport(proxyserver string) *http.Transport {
	transport := &http.Transport{}
	if proxyserver == "" {

		transport = &http.Transport{
			//!!!disable certificate CHECK!!!!
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		proxyStr := proxyserver
		proxyURL, err := url.Parse(proxyStr)
		transport = &http.Transport{
			//!!!disable certificate CHECK!!!!
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxyURL),
		}

		if err != nil {
			log.Println(err)
		}

	}
	return transport
}
