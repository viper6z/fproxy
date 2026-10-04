package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var proxyClient = newClient()

func newClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableCompression = true
	client := &http.Client{Transport: t}
	return client
}

func main() {
	proxyHandler := func(w http.ResponseWriter, req *http.Request) {
		out, err := http.NewRequestWithContext(req.Context(), req.Method, "http://127.0.0.1:9091"+req.URL.RequestURI(), req.Body)
		if err != nil {
			return
		}

		// Loop over header names
		for name, values := range req.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}

		for key, values := range req.Header {
			for _, v := range values {
				out.Header.Add(key, v)
			}
		}
		out.ContentLength = req.ContentLength

		resp, err := proxyClient.Do(out)
		if err != nil {
			return
		}

		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}

		w.WriteHeader(resp.StatusCode)

		io.Copy(w, resp.Body)
	}
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(proxyHandler)))
}
