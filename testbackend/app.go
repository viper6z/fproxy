package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const name string = "testhandler"

func main() {
	echoHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "backend: %s\n", name)
		fmt.Fprintf(w, "%s %s\n", req.Method, req.URL.String())
		fmt.Fprintf(w, "host: %s\n", req.Host)
		fmt.Fprintf(w, "remote: %s\n", req.RemoteAddr)
		for key, values := range req.Header {
			for _, v := range values {
				fmt.Fprintf(w, "%s: %s\n", key, v)
			}
		}

		fmt.Fprintf(w, "body: ")
		io.Copy(w, req.Body)
		fmt.Fprintf(w, "\n")
	}

	log.Fatal(http.ListenAndServe(":9091", http.HandlerFunc(echoHandler)))
}
