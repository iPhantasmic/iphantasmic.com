package main

import (
	"log"
	"net/http"

	"iphantasmic.com/internal/config"
	"iphantasmic.com/internal/handlers"
	"iphantasmic.com/internal/posts"
)

func main() {
	cfg := config.FromEnv()

	store, err := posts.LoadDir("internal/posts")
	if err != nil {
		log.Fatalf("load posts: %v", err)
	}

	server := handlers.NewServer(cfg, store)

	log.Printf("listening on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, server.Routes()); err != nil {
		log.Fatal(err)
	}
}
