---
title: "Hello GoTTH"
slug: "hello-gotth"
description: "A first markdown post rendered by the standalone Go proof of concept."
featured: true
published: "2026-06-01"
tags: ["go", "markdown", "poc"]
---

This is the first post in the standalone blog proof of concept.

It is intentionally small, but it already proves the important loop:

- load a markdown file from the repo
- parse frontmatter
- render the page through server-side Go templates
- serve CSS and JavaScript from `/assets`

## Why start with markdown?

Markdown is enough for the first version because the immediate goal is a blog.
We can still evolve the content model later if we want Notion-like blocks again.

> [!NOTE]
> The first milestone is ownership: content, rendering, and routing all live in the Go app.

## Code sample

```go
http.HandleFunc("/", handler)
```

This POC keeps the visual direction close to the current site: quiet document layout,
soft cards, expressive link accents, and a simple dark mode toggle.
