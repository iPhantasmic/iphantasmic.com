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
icon: "icon.svg"
cover: "cover.png"
canonical: "https://example.com/canonical-hello"
featured: true
published: "2026-06-01"
updated: "2026-06-04"
tags: ["go", "markdown", "go"]
---

## Heading

| Name | Value |
| --- | --- |
| Stack | GoTTH |

> [!WARNING]
> Relative media should resolve to the post content folder.

![Diagram](diagram.png?size=2)
![Root](/static/images/error.png)
`)
	writePost(t, dir, "later.md", `---
title: "Later"
slug: "later"
description: "A later normal post."
published: "2026-06-05"
---

Later body.
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
	writePost(t, dir, "whoami.md", `---
title: "whoami.exe"
slug: "whoami.exe"
kind: "page"
description: "Profile page."
published: "2026-06-03"
---

Profile page.
`)

	store, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	all := store.All()
	if len(all) != 2 {
		t.Fatalf("len(All()) = %d, want 2", len(all))
	}

	post := all[0]
	if post.Title != "Hello" {
		t.Fatalf("Title = %q, want featured Hello first", post.Title)
	}
	if got := strings.Join(post.Tags, ","); got != "go,markdown" {
		t.Fatalf("Tags = %q, want go,markdown", got)
	}
	if post.Icon != "/static/content/hello/icon.svg" {
		t.Fatalf("Icon = %q, want /static/content/hello/icon.svg", post.Icon)
	}
	if post.Cover != "/static/content/hello/cover.png" {
		t.Fatalf("Cover = %q, want /static/content/hello/cover.png", post.Cover)
	}
	if post.Canonical != "https://example.com/canonical-hello" {
		t.Fatalf("Canonical = %q, want https://example.com/canonical-hello", post.Canonical)
	}
	if !post.Featured {
		t.Fatal("Featured = false, want true")
	}
	if got := post.Updated.Format("2006-01-02"); got != "2026-06-04" {
		t.Fatalf("Updated = %q, want 2026-06-04", got)
	}
	if got := post.LastModified().Format("2006-01-02"); got != "2026-06-04" {
		t.Fatalf("LastModified() = %q, want 2026-06-04", got)
	}
	if got := post.URLPath(); got != "/posts/hello" {
		t.Fatalf("URLPath() = %q, want /posts/hello", got)
	}
	if !strings.Contains(string(post.Body), "<table>") {
		t.Fatalf("Body did not render a GFM table: %s", post.Body)
	}
	if !strings.Contains(string(post.Body), `class="callout callout-warning"`) {
		t.Fatalf("Body did not render a warning callout: %s", post.Body)
	}
	if strings.Contains(string(post.Body), "[!WARNING]") {
		t.Fatalf("Body still includes callout marker: %s", post.Body)
	}
	if !strings.Contains(string(post.Body), `src="/static/content/hello/diagram.png?size=2"`) {
		t.Fatalf("Body did not resolve a relative image path: %s", post.Body)
	}
	if !strings.Contains(string(post.Body), `src="/static/images/error.png"`) {
		t.Fatalf("Body rewrote a root-relative image path: %s", post.Body)
	}
	if _, ok := store.Find("draft"); ok {
		t.Fatal("Find(draft) found a draft post")
	}
	if page, ok := store.Find("whoami.exe"); !ok || page.Kind != "page" || page.URLPath() != "/whoami.exe" {
		t.Fatalf("Find(whoami.exe) = (%+v, %v), want page", page, ok)
	}
	if got := len(store.SitemapPages()); got != 3 {
		t.Fatalf("len(SitemapPages()) = %d, want 3", got)
	}
}

func TestLoadDirValidatesUpdatedDate(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "bad.md", `---
title: "Bad"
slug: "bad"
description: "Invalid updated date."
published: "2026-06-01"
updated: "June 4, 2026"
---

Body.
`)

	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("LoadDir() error = nil, want updated date validation error")
	}
	if !strings.Contains(err.Error(), "updated must use YYYY-MM-DD") {
		t.Fatalf("LoadDir() error = %v, want updated date validation error", err)
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
