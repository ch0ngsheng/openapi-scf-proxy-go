package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// New takes target host and creates a reverse proxy
func New(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req, url.Host)
	}

	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request, targetHost string) {
	req.Host = targetHost
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41")

	req.Header.Del("x-forwarded-for")
	req.Header.Del("x-real-ip")
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		if err != nil {
			fmt.Printf("Got error while modifying response: %v \n", err)
		}
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		resp.Header.Set("x-proxy-by", "go")
		log.Println(resp.StatusCode)
		return nil
	}
}

// RequestHandler handles the http request using proxy
func RequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
