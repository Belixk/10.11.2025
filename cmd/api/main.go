package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Belixk/10.11.2025/internal/handlers"
	"github.com/Belixk/10.11.2025/internal/storage"
)

const (
	port = ":8080"
)

func main() {
	go func() {
		for range time.Tick(25 * time.Second) {
			storage.LoadFromFile()
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		<-ctx.Done()
		storage.SaveToFile()
		stop()
		os.Exit(0)
	}()

	http.HandleFunc("/check_links", handlers.HandleCheckLinks)
	http.HandleFunc("/report", handlers.HandleReport)

	fmt.Println("===Запуск сервера===")
	http.ListenAndServe(port, nil)
}
