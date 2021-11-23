package main

import (
	"context"
	"github.com/SuperMohit/orchestrator/internal/manifests"
	"github.com/SuperMohit/orchestrator/internal/routes"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	maxHeaderBytes    = 1024
	readHeaderTimeout = 0o3
	readTimeout       = 0o3
	writeTimeout      = 15
	idleTimeout       = 60
	graceTime         = 15
)

func main()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Info("started the server")

	composer := manifests.NewDockerComposer()

	router := routes.NewOrchestrateRouter(composer)

	listen(":9001", router.ServiceRouter())
}


func listen(address string, handler http.Handler) {
	server := &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadTimeout:       readTimeout * time.Second,
		ReadHeaderTimeout: readHeaderTimeout * time.Second,
		WriteTimeout:      writeTimeout * time.Second,
		IdleTimeout:       idleTimeout * time.Second,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Println("Started and Listening at address: ", address)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Error and Shutting down")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), graceTime*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error resulted in shutdown.")
	}
}