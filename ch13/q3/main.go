package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type timeResponse struct {
	DayWeek  string `json:"day_of_week"`
	DayMonth int    `json:"day_of_month"`
	Month    string `json:"month"`
	Year     int    `json:"year"`
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
	Second   int    `json:"second"`
}

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
		now := time.Now()
		if r.Header.Get("Accept") == "application/json" {
			response := timeResponse{
				DayWeek:  now.Weekday().String(),
				DayMonth: now.Day(),
				Month:    now.Month().String(),
				Year:     now.Year(),
				Hour:     now.Hour(),
				Minute:   now.Minute(),
				Second:   now.Second(),
			}
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(200)
			encoder := json.NewEncoder(w)
			encoder.Encode(response)
		} else {
			rfc3339 := now.Format(time.RFC3339)
			w.Header().Add("content-type", "text/plain")
			w.WriteHeader(200)
			w.Write([]byte("The current time is: " + rfc3339))
		}
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
