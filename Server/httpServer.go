package main

import (
	"Server/internal"
	"log"
	"net/http"
	"net/http/pprof"
	//_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
)

func main() {
	r := http.NewServeMux()

	//Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	server := internal.NewServer()
	server.Run()
	http.HandleFunc("/get", server.Get)
	http.HandleFunc("/set", server.Set)

	r.HandleFunc("/get", server.Get)
	r.HandleFunc("/set", server.Set)

	runtime.SetBlockProfileRate(-1)

	portNumber := 8000
	if len(os.Args) > 1 {
		portNumber, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portNumber)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portNumber), r))
	//log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portNumber), nil))
}
