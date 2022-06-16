package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultPort = 8080
)

func main() {
	// Return a server
	srv := http.Server{}

	// Create a channel that will be used to listen for cancellation signales
	done := make(chan os.Signal, 1)

	// Relay signals to the channel
	// We listen for signals, capture them if they happen and are then sent to the channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("an error occured while starting the server: %v", err)
		}
	}()
	logrus.Printf("server running on port: %v", defaultPort)

	// This listens for output from the `done` channel created.
	// If there's any output, it'll begin the process of shutting down the server
	<-done

	cancelContext, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer func() {
		// closes databases, etc
		cancel()
	}()

	err := srv.Shutdown(cancelContext)
	logrus.Print("server shutting down ...")
	if err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
	logrus.Print("server exited properly")
}
