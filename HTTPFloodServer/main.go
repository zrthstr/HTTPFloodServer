package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	var showHelp bool

	// Define command line flags
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.BoolVar(&showHelp, "h", false, "Show help message")

	// Parse command line flags
	flag.Parse()

	// Show help message if -h flag is present
	if showHelp {
		flag.Usage()
		return
	}

	// Create a new HTTP server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Define a handler function that sends an infinite stream of data to the client
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers to indicate that the response is an infinite stream of data
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Transfer-Encoding", "chunked")

		// Create a buffer of 64KB to use for sending data
		//data := make([]byte, 64*1024)
		data := make([]byte, 1000*1000)
		stat := 0

		// Loop forever, sending the data buffer to the client
		for {
			_, err := w.Write(data)
			stat++
			if err != nil {
				log.Printf("Error sending data: %v", err)
				return
			}

			// Print transmission stats every 1 second
			//time.Sleep(1 * time.Second)
			log.Printf("Transmitted %d MB", stat)
		}
	})

	// Start the HTTP server
	log.Printf("Listening on port %d...", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
