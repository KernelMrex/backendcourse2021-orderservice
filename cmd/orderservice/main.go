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
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	serverUrl := ":8080"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	srv := startServer(serverUrl)
	waitForKillSignal(getKillSignalChan())
	srv.Shutdown(context.Background())
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
