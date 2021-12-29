package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting the grpc server at Port: 5000")
	log.Println("Starting the http server at Port: 5050")

	ss, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Failed to start server.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to user service. Thank you")
	})

	s := &http.Server{
		Addr:           ":5050",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var httpServerError = s.ListenAndServe()

	if httpServerError != nil {
		log.Fatalf("Failed to start http server.")
	}

	ser := grpc.NewServer()

	if err := ser.Serve(ss); err != nil {
		log.Fatalf("Failed to serve", err)
	}

}
