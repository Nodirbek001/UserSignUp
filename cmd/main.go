package main

import (
	"NewProUser/ci"
	"NewProUser/configs"
	"NewProUser/pkg/middleware"
	"NewProUser/service"

	"github.com/blendle/zapdriver"
	"github.com/gorilla/mux"
)

func main() {
	conf := configs.Load()
	if err := conf.Validate(); err != nil {
		panic(err)

	}
	logger, err := zapdriver.NewDevelopment()

	if err != nil {
		panic(err)
	}

	ci.MigrationsUp()
	userService := service.NewUserService(logger)

	root := mux.NewRouter()

	root.Use(middleware.PanicRecovery)
	root.Use(middleware.Logging)
}
