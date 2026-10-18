package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func newProxy(target *url.URL) *httputil.ReverseProxy {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.Proxy = nil
	t.MaxIdleConns = 200
	t.MaxConnsPerHost = 100
	t.ResponseHeaderTimeout = 10 * time.Second
	// TODO: add tls config
	return &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(target)
		},
	}
}

const (
	certfile = "certs/server.crt"
	keyfile  = "certs/server.key"
)

func main() {
	backendAddress, err := url.Parse("http://127.0.0.1:9091")
	if err != nil {
		log.Fatalf("error parsing listener url: %v", err)
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        newProxy(backendAddress),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServeTLS(certfile, keyfile))
}
