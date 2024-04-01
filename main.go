package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/viktoriina/databases-lab1/internal/app"
	"github.com/viktoriina/databases-lab1/internal/config"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	a := app.NewApp(cfg)
	defer func() {
		if err := a.Shutdown(); err != nil {
			log.Fatalf("Application shutdown error: %v", err)
		}
	}()
	done := make(chan error, 1)
	go func() {
		done <- a.Start()
	}()

	select {
	case <-signals:
		fmt.Println()
		log.Println("Received OS signal, shutting down gracefully..")
	case err := <-done:
		if err != nil {
			log.Fatalf("Application error: %v", err)
		}
		fmt.Println()
		log.Println("Application exited successfully.")
	}
}
