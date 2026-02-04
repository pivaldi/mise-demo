package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	port := getenv("API_PORT", "8080")
	dsn := getenv("DATABASE_URL", "postgres://app:app@127.0.0.1:5438/app?sslmode=disable")

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		conn, err := pgx.Connect(ctx, dsn)
		if err != nil {
			http.Error(w, "db connect: "+err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer conn.Close(ctx)

		rows, err := conn.Query(ctx, "select id, name, created_at from items order by id asc")
		if err != nil {
			http.Error(w, "query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		items := []Item{}
		for rows.Next() {
			var it Item
			if err := rows.Scan(&it.ID, &it.Name, &it.CreatedAt); err != nil {
				http.Error(w, "scan: "+err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, it)
		}

		setHeader(w)
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		_ = enc.Encode(map[string]any{
			"dsn_hint":    "DATABASE_URL via mise env",
			"items":       items,
			"items_count": len(items),
		})
	})

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           logging(mux),
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Printf("[api] listening on http://localhost:%s", port)
	log.Fatal(srv.ListenAndServe())
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[api] %s %s (%s)", r.Method, r.URL.Path, time.Since(start).Truncate(time.Millisecond))
	})
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}
