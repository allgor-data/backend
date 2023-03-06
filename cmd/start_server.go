package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/allgor-data/backend/app/api/graphql/generated"
	"github.com/allgor-data/backend/app/api/graphql/resolver"
	"github.com/rs/zerolog/log"
)

func startServer(port string, enablePlayground bool) error {
	svMux := http.NewServeMux()

	gqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver.Resolver{},
			},
		),
	)

	svMux.Handle("/", gqlHandler)
	log.Info().Msg("Added GraphQL handler at /")

	if enablePlayground {
		svMux.Handle("/playground", playground.Handler("playground", "/"))
		log.Info().Msg("Added playground handler at /playground")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: svMux,
	}

	go func() {
		log.Info().Msgf("Server listening at http://localhost:%s/", port)

		err := server.ListenAndServe()
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			log.Info().Msg("Stopped serving new connections")
		} else {
			log.Err(err).Msg("HTTP server error")
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Err(err).Msg("HTTP shutdown error")
	}

	log.Info().Msg("Graceful shutdown complete")

	return nil
}
