package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"net/http"
)

type TestServer struct {
}

func (ts TestServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (ts TestServer) Serve(certFile, keyFile string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { ts.handleRoot(w, r) })
	server := http.Server{}
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Error().Msgf("Error when trying to ListenAndServe: %v", err)
	}
}

func main() {
	crtOption := flag.String("crt", "", "")
	keyOption := flag.String("key", "", "")
	flag.Parse()

	if crtOption == nil || keyOption == nil {
		log.Fatal().Msgf("Provide crt and key")
	}

	ts := TestServer{}
	ts.Serve(*crtOption, *keyOption)
}
