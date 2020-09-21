package main

import (
	"fmt"
	"github.com/dalmarcogd/gbpl-go/internal/cache"
	"github.com/dalmarcogd/gbpl-go/internal/database"
	"github.com/dalmarcogd/gbpl-go/internal/environment"
	"github.com/dalmarcogd/gbpl-go/internal/grpcserver"
	"github.com/dalmarcogd/gbpl-go/internal/handlers"
	"github.com/dalmarcogd/gbpl-go/internal/logger"
	"github.com/dalmarcogd/gbpl-go/internal/services"
	"os"
	"os/signal"
)

func main() {
	ss := services.
		New().
		WithDatabase(database.New()).
		WithCache(cache.New()).
		WithLogger(logger.New()).
		WithGrpcServer(grpcserver.New().WithAddress(":8080")).
		WithHandlers(handlers.New()).
		WithEnvironment(environment.New())

	if err := ss.Init(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}

	go func() {
		ss.Logger().Info(ss.Context(), "Grpc server started")
		if err := ss.GrpcServer().Run(); err != nil {
			ss.Logger().Fatal(ss.Context(), err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	sig := <-quit

	ss.Logger().Info(ss.Context(), fmt.Sprintf("Shutdown by %v", sig.String()))

	if err := ss.Close(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}
	ss.Logger().Info(ss.Context(), "All services closed")
}
