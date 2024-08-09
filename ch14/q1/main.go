package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	timeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wait := rand.Intn(5000)

		select {
		case <-time.After(time.Millisecond * time.Duration(wait)):
			time := time.Now().Format(time.RFC3339)
			w.Header().Add("content-type", "text/plain")
			w.WriteHeader(200)
			w.Write([]byte("The current time is: " + time))
			return
		case <-r.Context().Done():
			http.Error(w, "request timeout", http.StatusRequestTimeout)
			return
		}

	})
	timeoutMw := timeoutMiddleware(2000)

	mux := http.NewServeMux()
	mux.Handle("GET /time", timeoutMw(timeHandler))

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

func timeoutMiddleware(timeout int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), time.Duration(time.Millisecond*time.Duration(timeout)))
			defer cancel()
			reqWithContext := r.WithContext(ctx)
			h.ServeHTTP(w, reqWithContext)
		})
	}
}
