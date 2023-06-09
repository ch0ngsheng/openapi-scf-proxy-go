package proxy

import (
	"context"
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
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	req.Header["X-Forwarded-For"] = nil
	req.Header.Del("X-Real-Ip")
	req.Header.Del("X-Forwarded-Proto")
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

// RequestHandler2 handles the http request using proxy for aliyun
func RequestHandler2(proxy *httputil.ReverseProxy) func(context.Context, http.ResponseWriter, *http.Request) {
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
