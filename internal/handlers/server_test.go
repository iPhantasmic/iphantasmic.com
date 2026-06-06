package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"iphantasmic.com/internal/config"
	"iphantasmic.com/internal/posts"
)

func TestRoutesServeMarkdownPagesBySlug(t *testing.T) {
	dir := t.TempDir()
	writeMarkdown(t, dir, "about.md", `---
title: "About"
slug: "about"
kind: "page"
description: "About page."
published: "2026-06-06"
---

About body.
`)

	store, err := posts.LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	server := NewServer(config.Site{
		Name:        "Test Site",
		Author:      "Tester",
		Description: "Test description.",
		BaseURL:     "https://example.com",
	}, store)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /about status = %d, want %d", rec.Code, http.StatusOK)
	}
	if !strings.Contains(rec.Body.String(), "About body.") {
		t.Fatalf("GET /about body did not include page content: %s", rec.Body.String())
	}

	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/posts/about", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("GET /posts/about status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func writeMarkdown(t *testing.T, dir, name, body string) {
	t.Helper()

	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", name, err)
	}
}
