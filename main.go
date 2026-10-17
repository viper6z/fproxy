package main

import (
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

func newHandler(client *http.Client, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		out, err := http.NewRequestWithContext(req.Context(), req.Method, target+req.URL.RequestURI(), req.Body)
		if err != nil {
			return
		}

		for key, values := range req.Header {
			for _, v := range values {
				out.Header.Add(key, v)
			}
		}
		out.ContentLength = req.ContentLength

		resp, err := client.Do(out)
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
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", newHandler(proxyClient, "http://127.0.0.1:9091")))
}
