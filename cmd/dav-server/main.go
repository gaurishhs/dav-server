package main

import (
	"os"

	"github.com/gaurishhs/dav-server/internal/server"
	"github.com/ory/graceful"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pflag "github.com/spf13/pflag"
)

func main() {
	var (
		debug   bool
		logJSON bool
	)

	pflag.BoolVarP(&debug, "debug", "d", false, "toggle debug logs")
	pflag.BoolVarP(&logJSON, "json", "j", true, "toggle json logs")
	pflag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if !logJSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	davServer, err := server.NewDAVServer()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create a new dav server")
	}
	server := graceful.WithDefaults(davServer.HttpServer)

	log.Info().Msgf("starting server at addr: %s", server.Addr)
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown")
	}
	log.Info().Msg("server was shutdown gracefully")
}
