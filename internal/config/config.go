package config

import (
	"os"
	"strings"
)

type Site struct {
	Name        string
	Author      string
	Description string
	Domain      string
	BaseURL     string
	Addr        string
}

func FromEnv() Site {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "https://iphantasmic.com"
	}
	baseURL = strings.TrimRight(baseURL, "/")

	return Site{
		Name:        "iPhantasmic's Blog",
		Author:      "iPhantasmic",
		Description: "A standalone Go markdown blog POC.",
		Domain:      "iphantasmic.com",
		BaseURL:     baseURL,
		Addr:        addr,
	}
}

func (s Site) URL(path string) string {
	if path == "" || path == "/" {
		return s.BaseURL
	}
	if path[0] != '/' {
		path = "/" + path
	}
	return s.BaseURL + path
}
