package main

import (
	"log"

	"github.com/jgoerz/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
