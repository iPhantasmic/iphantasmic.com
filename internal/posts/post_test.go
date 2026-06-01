package posts

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadDirLoadsPublishedPosts(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "hello.md", `---
title: "Hello"
slug: "hello"
description: "A real post."
published: "2026-06-01"
tags: ["go", "markdown", "go"]
---

## Heading

| Name | Value |
| --- | --- |
| Stack | GoTTH |
`)
	writePost(t, dir, "draft.md", `---
title: "Draft"
slug: "draft"
description: "Hidden."
published: "2026-06-02"
draft: true
---

This should not load.
`)

	store, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	all := store.All()
	if len(all) != 1 {
		t.Fatalf("len(All()) = %d, want 1", len(all))
	}

	post := all[0]
	if post.Title != "Hello" {
		t.Fatalf("Title = %q, want Hello", post.Title)
	}
	if got := strings.Join(post.Tags, ","); got != "go,markdown" {
		t.Fatalf("Tags = %q, want go,markdown", got)
	}
	if !strings.Contains(string(post.Body), "<table>") {
		t.Fatalf("Body did not render a GFM table: %s", post.Body)
	}
	if _, ok := store.Find("draft"); ok {
		t.Fatal("Find(draft) found a draft post")
	}
}

func TestLoadDirRejectsDuplicateSlugs(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "one.md", postWithSlug("same"))
	writePost(t, dir, "two.md", postWithSlug("same"))

	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("LoadDir() error = nil, want duplicate slug error")
	}
	if !strings.Contains(err.Error(), "duplicate post slug") {
		t.Fatalf("LoadDir() error = %v, want duplicate slug error", err)
	}
}

func TestLoadDirValidatesFrontMatter(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "bad.md", `---
title: "Bad"
slug: "Bad Slug"
description: "Invalid slug."
published: "2026-06-01"
---

Body.
`)

	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("LoadDir() error = nil, want validation error")
	}
	if !strings.Contains(err.Error(), "lowercase kebab-case") {
		t.Fatalf("LoadDir() error = %v, want slug validation error", err)
	}
}

func writePost(t *testing.T, dir, name, body string) {
	t.Helper()

	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", name, err)
	}
}

func postWithSlug(slug string) string {
	return `---
title: "Post"
slug: "` + slug + `"
description: "Description."
published: "2026-06-01"
---

Body.
`
}
