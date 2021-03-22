package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"orderservice/pkg/orderservice/transport"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	cfg, err := parseEnv()
	if err != nil {
		log.Fatal("error parsing env vars")
		return
	}

	log.WithFields(log.Fields{"url": cfg.ServeRESTAddress}).Info("starting the server...")
	srv := startServer(cfg.ServeRESTAddress)
	waitForKillSignal(getKillSignalChan())
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func startServer(serverUrl string) *http.Server {
	router := transport.Router()
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatalln(srv.ListenAndServe())
	}()
	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
