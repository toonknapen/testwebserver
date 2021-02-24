package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/toonknapen/testwebserver/internal/server"
)

func main() {
	crtOption := flag.String("crt", "", "")
	keyOption := flag.String("key", "", "")
	flag.Parse()

	if len(*crtOption) == 0 || len(*keyOption) == 0 {
		log.Fatal().Msgf("Provide crt and key")
	}

	ts := server.TestServer{}
	ts.Serve(*crtOption, *keyOption)
}
