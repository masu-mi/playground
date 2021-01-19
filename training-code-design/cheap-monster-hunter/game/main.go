package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/controller/http"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/gateway/dummy"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain/service"
)

func init() {
	service.Logger = log.New(os.Stdout, "[chep-monster-hunter:http]: ", log.Lshortfile)
}

func main() {
	handler := app.NewHTTPHandler(&dummy.Gateway{
		HunterRepo:  &dummy.HunterRepo{},
		MonsterRepo: &dummy.MonsterRepo{},
	})
	handler.SetLogger(service.Logger)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
