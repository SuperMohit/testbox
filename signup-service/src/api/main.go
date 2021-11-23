package main

import (
	"context"
	"github.com/SuperMohit/signup/internal"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := dbConnection()
	if err !=nil {
		return
	}

	repo := internal.NewSignupRepo(db)
	//repo.Migration()
	router := internal.NewSignupRouter(repo)
	listen(":9002", router.ServiceRouter())
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

func dbConnection() (*gorm.DB, error) {
	dsn := "host=testingmohit.cregpjltzpr2.ap-south-1.rds.amazonaws.com user=postgres password=51rWulxAIJ6G8doXuPZL dbname=account port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil,err
	}
	return db, nil
}