package config

import "os"

type Site struct {
	Name        string
	Author      string
	Description string
	Domain      string
	Addr        string
}

func FromEnv() Site {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return Site{
		Name:        "iPhantasmic's Blog",
		Author:      "iPhantasmic",
		Description: "A standalone Go markdown blog POC.",
		Domain:      "iphantasmic.com",
		Addr:        addr,
	}
}
