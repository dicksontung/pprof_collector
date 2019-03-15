package server

import (
	"log"
	"net/http"
	"net/http/pprof"
)

//Profiling enables the pprof server
func Profiling() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(":1303", mux)
	if err != nil {
		log.Printf("Unable to start pprof endpoint: %+v", err)
	}
}
