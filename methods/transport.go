package createtransport

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func CreateTransport(proxyserver string) *http.Transport {
	transport := &http.Transport{}
	if proxyserver == "" {

		transport = &http.Transport{

			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
	} else {
		proxyStr := proxyserver
		fmt.Println("!!Certificate check disabled!!")
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
