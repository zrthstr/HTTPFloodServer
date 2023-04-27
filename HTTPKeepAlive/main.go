package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "port number")
	help := flag.Bool("h", false, "help")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	i := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Connection", "keep-alive")

		// Write the response
		w.Write([]byte("Hello, world!"))
		i++
		fmt.Printf("connection: %d", i)

		// Keep the connection open forever by preventing the response writer
		// from being closed
		select {}
	})

	// Listen and serve on port 8080
	addr := fmt.Sprintf(":%d", *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
