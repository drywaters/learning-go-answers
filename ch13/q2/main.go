package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stdout, options)
	slogger := slog.New(handler)
	logMiddlware := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			slogger.Info("incoming request", "ip", r.RemoteAddr)
			h.ServeHTTP(w, r)
		}
	}
	timeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().Format(time.RFC3339)
		w.Header().Add("content-type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("The current time is: " + time))
	})

	mux := http.NewServeMux()
	mux.Handle("GET /time", logMiddlware(timeHandler))

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	fmt.Println("Listening on port: 8080")
	server.ListenAndServe()
}
