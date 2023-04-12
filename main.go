package main

import (
	"NewProUser/api"
	"NewProUser/ci"
	"NewProUser/configs"
	"NewProUser/pkg/middleware"
	"NewProUser/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/blendle/zapdriver"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
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


	casbinJWTRoleAuthorizer, err := middleware.NewCasbinJWTRoleAuthorizer(&conf)
	if err != nil {
		logger.Fatal("Could not initialize Cabin JWT Role Authorizer", zap.Error(err))
	}


	root.Use(casbinJWTRoleAuthorizer.Middleware)
	api.Init(root, userService, logger)
	log.Println("main:Project is started on the Port", conf.HTTPPort)


	errchan:=make(chan error, 1)
	osSignals:=make(chan os.Signal,1)


	signal.Notify(osSignals,os.Interrupt, syscall.SIGTERM)

	httpServer:=http.Server{
		Addr: conf.HTTPPort,
		Handler: root,
	}


	// http server
	go func() {
		errchan <- httpServer.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-errchan:

		logger.Fatal("error: ", zap.Error(err))

	case <-osSignals:
		logger.Info("main : received os signal, shutting down")
		_ = httpServer.Shutdown(context.Background())
		return
	}


}
