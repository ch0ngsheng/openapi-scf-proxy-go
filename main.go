package main

import (
	"log"
	"net/http"
	"os"

	pp "openapi-scf-proxy-go/proxy"
)

var (
	Target   string
	ListenOn string
)

func init() {
	// TargetHostURL=https://api.openai.com
	Target = os.Getenv("TargetHostURL")
	// ProxyListenOn=0.0.0.0:9000
	ListenOn = os.Getenv("ProxyListenOn")
}

func main() {
	if len(Target) == 0 || len(ListenOn) == 0 {
		panic("ProxyHostURL or ProxyListenOn is null.")
	}

	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := pp.New(Target)
	if err != nil {
		panic(err)
	}

	// handle all requests to your server using the proxy
	http.HandleFunc("/", pp.RequestHandler(proxy))
	log.Fatal(http.ListenAndServe(ListenOn, nil))
}
