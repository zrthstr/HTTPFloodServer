package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Define command line flags
	port := flag.Int("p", 8080, "Port number to listen on")
	bufferSize := flag.Int("bufferSize", 1024*1024*1024, "Size of buffer to generate in bytes")
	chunkSize := flag.Int("chunkSize", 1024*1024, "Size of each chunk to send in bytes")
	compress := flag.Bool("compress", true, "Whether to compress the buffer")
	flag.Parse()

	// Generate the buffer
	buffer := bytes.NewBuffer(make([]byte, 0, *bufferSize))
	for i := 0; i < *bufferSize; i++ {
		buffer.WriteByte(byte(i % 256))
	}

	// Compress the buffer if requested
	if *compress {
		var compressed bytes.Buffer
		gzipWriter := gzip.NewWriter(&compressed)
		if _, err := io.Copy(gzipWriter, buffer); err != nil {
			log.Fatal(err)
		}
		if err := gzipWriter.Close(); err != nil {
			log.Fatal(err)
		}
		buffer = &compressed
	}

	// Define the HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Length header
		w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))

		// Set the Content-Encoding header if compressed
		if *compress {
			w.Header().Set("Content-Encoding", "gzip")
		}

		// Send the buffer in chunks
		for {
			chunk := make([]byte, *chunkSize)
			n, err := buffer.Read(chunk)
			if err != nil && err != io.EOF {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if n == 0 {
				break
			}
			if _, err := w.Write(chunk[:n]); err != nil {
				log.Println(err)
				return
			}
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}

	// Create the HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.HandlerFunc(handler),
	}

	// Start the HTTP server
	log.Printf("Listening on port %d...\n", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
