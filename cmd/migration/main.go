package main

import (
	"github.com/dalmarcogd/gbpl-go/internal/database"
	"github.com/dalmarcogd/gbpl-go/internal/environment"
	"github.com/dalmarcogd/gbpl-go/internal/logger"
	"github.com/dalmarcogd/gbpl-go/internal/models"
	"github.com/dalmarcogd/gbpl-go/internal/services"
)

func main() {
	ss := services.
		New().
		WithDatabase(database.New()).
		WithLogger(logger.New()).
		WithEnvironment(environment.New())

	if err := ss.Init(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}

	if err := ss.Database().DB(ss.Context()).AutoMigrate(&models.User{}); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}
	ss.Logger().Info(ss.Context(), "Migration finished")
}
