package main

import (
	"common"
	"log"
	"net/http"

	"errors"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)
	log.Printf("Server listening on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server")
	}

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-term
		if err := srv.Close(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error closing Server: %v", err)
		}
	}()

}
