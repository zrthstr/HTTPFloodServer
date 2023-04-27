package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("p", 8080, "port number")
	help := flag.Bool("h", false, "help")
	numHeaders := flag.Int("n", 10, "number of headers to send per request")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	count := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Generate and send the specified number of random HTTP headers
		for i := 0; i < *numHeaders; i++ {
			// Generate a random HTTP header
			headerName := "Header-" + strconv.Itoa(rand.Intn(1000000))
			headerValue := strconv.Itoa(rand.Intn(100000))
			w.Header().Set(headerName, headerValue)

			// Increment the counter and print stats every 100 headers
			count++
			if count%100 == 0 {
				fmt.Printf("%d headers sent\n", count)
			}
		}

		// Send a small payload after the headers
		fmt.Fprintln(w, "Hello, world!")
	})

	// Start the server on the specified port
	fmt.Printf("Listening on :%d...\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic(err)
	}
}
