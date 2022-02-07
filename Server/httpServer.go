package main

import (
	"Server/internal"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	server := internal.NewServer()
	server.Run()
	http.HandleFunc("/get", server.Get)
	http.HandleFunc("/set", server.Set)

	portNumber := 8000
	if len(os.Args) > 1 {
		portNumber, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portNumber)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portNumber), nil))
}
