package server

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
	"time"
)

type TestServer struct {
}

func (ts TestServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	queryParamMap := r.URL.Query()
	responseTime, ok := queryParamMap["responsetime"]
	if ok {
		waitTime, err := strconv.Atoi(responseTime[0])
		if err != nil {
			log.Info().Msgf("Could not convert responsetime arg to int: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		time.Sleep(time.Duration(waitTime) * time.Second)
	}

	WriteJSON(w, *r)
}

func (ts TestServer) Serve(certFile, keyFile string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { ts.handleRoot(w, r) })

	server := http.Server{
		Addr:    ":7999",
		Handler: mux,
	}
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Error().Msgf("Error when trying to ListenAndServe: %v", err)
	}
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Error().Msgf("io.Copy: %v", err)
		return
	}
}
