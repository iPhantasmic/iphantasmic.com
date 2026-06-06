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

func TestSearchRoute(t *testing.T) {
	dir := t.TempDir()
	writeMarkdown(t, dir, "red-team.md", `---
title: "Red Team Notes"
slug: "red-team-notes"
description: "Adversary simulation notes."
published: "2026-06-06"
tags: ["security"]
---

Payload design and detection notes.
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
	req := httptest.NewRequest(http.MethodGet, "/search?q=payload", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /search status = %d, want %d", rec.Code, http.StatusOK)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "Red Team Notes") {
		t.Fatalf("GET /search did not include result title: %s", body)
	}
	if !strings.Contains(body, "/posts/red-team-notes") {
		t.Fatalf("GET /search did not include result link: %s", body)
	}

	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/search?q=payload&partial=1", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /search partial status = %d, want %d", rec.Code, http.StatusOK)
	}
	body = rec.Body.String()
	if !strings.Contains(body, `data-search-results`) {
		t.Fatalf("GET /search partial did not include search results container: %s", body)
	}
	if strings.Contains(body, "<html") {
		t.Fatalf("GET /search partial returned a full document: %s", body)
	}
	if !strings.Contains(body, "Red Team Notes") {
		t.Fatalf("GET /search partial did not include result title: %s", body)
	}
}

func TestTagRoutes(t *testing.T) {
	dir := t.TempDir()
	writeMarkdown(t, dir, "red-team.md", `---
title: "Red Team Notes"
slug: "red-team-notes"
description: "Adversary simulation notes."
published: "2026-06-06"
tags: ["security", "red team"]
---

Payload design and detection notes.
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
	req := httptest.NewRequest(http.MethodGet, "/tags", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /tags status = %d, want %d", rec.Code, http.StatusOK)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "security") || !strings.Contains(body, "/tags/security") {
		t.Fatalf("GET /tags did not include security tag link: %s", body)
	}
	if !strings.Contains(body, "red team") || !strings.Contains(body, "/tags/red-team") {
		t.Fatalf("GET /tags did not include red team tag link: %s", body)
	}

	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/tags/security", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /tags/security status = %d, want %d", rec.Code, http.StatusOK)
	}
	body = rec.Body.String()
	if !strings.Contains(body, "Red Team Notes") || !strings.Contains(body, "/posts/red-team-notes") {
		t.Fatalf("GET /tags/security did not include tagged post: %s", body)
	}

	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/tags/unknown", nil)
	server.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("GET /tags/unknown status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func writeMarkdown(t *testing.T, dir, name, body string) {
	t.Helper()

	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", name, err)
	}
}
