package main

import (
	"context"
	"fmt"
	"net/http"
)

type levelKey struct{}
type Level string

const Debug = Level("debug")
const Info = Level("info")

func main() {
}

func LogLevelMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")

		if level != "debug" && level != "info" {
			http.Error(w, "invalid log type", http.StatusBadRequest)
			return
		}

		ctx := ContextWithLevel(r.Context(), Level(level))
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	level, ok := LevelFromContext(ctx)
	if !ok {
		// no level set, just return with no message
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, levelKey{}, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(levelKey{}).(Level)
	return level, ok
}
