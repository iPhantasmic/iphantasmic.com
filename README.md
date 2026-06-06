# iPhantasmic's Blog

Standalone GoTTH markdown blog for `iphantasmic.com`.

The former Next.js + Notion implementation is parked in [`legacy-next/`](./legacy-next)
for reference while this rewrite takes over the root project.

## Stack

- Go
- templ
- Tailwind CSS
- Markdown with YAML frontmatter
- Air for local live reload

HTMX is planned for interactive features once the core blog is settled.

## Quick Start

Install tools and dependencies:

```bash
make install-tools
```

Build CSS:

```bash
make css
```

Run the app:

```bash
make run
```

Or run with Air live reload:

```bash
make dev
```

The default local URL is:

```txt
http://localhost:8090
```

## Useful Commands

```bash
make templ       # generate templ Go files
make css         # build assets/css/app.css
make css-watch   # watch Tailwind input
make test        # generate templates and run Go tests
make build       # build ./tmp/server
make clean       # remove local build output
```

## Environment

```bash
ADDR=:8090
BASE_URL=http://localhost:8090
```

`BASE_URL` is used for RSS, sitemap, and robots output.

## Content

Markdown files live in [`internal/posts/`](./internal/posts).

Required frontmatter:

```yaml
---
title: "Hello GoTTH"
slug: "hello-gotth"
kind: "post"
description: "Short summary for lists, RSS, and metadata."
published: "2026-06-01"
tags: ["go", "markdown"]
---
```

Full frontmatter schema:

```yaml
---
title: "Hello GoTTH"              # required
slug: "hello-gotth"               # required, lowercase kebab-case or dot-separated
kind: "post"                      # optional, defaults to post
description: "Short summary."     # required
icon: "/static/images/icon.png"    # optional emoji or image URL/path
cover: "/static/images/cover.png" # optional social preview image URL/path
canonical: "https://example.com/original-post" # optional canonical override
published: "2026-06-01"           # required, YYYY-MM-DD
updated: "2026-06-04"             # optional, YYYY-MM-DD
featured: false                   # optional, reserved for future homepage treatment
tags: ["go", "markdown"]          # optional
draft: false                      # optional
---
```

Supported `kind` values:

- `post`: appears on the homepage and in `/feed.xml`
- `page`: addressable at `/{slug}` and included in `/sitemap.xml`, but hidden from the post list and feed

Drafts can be excluded with:

```yaml
draft: true
```

## Starter Pages

- `/posts/hello-gotth`
- `/whoami.exe`

The `whoami.exe` page is backed by [`internal/posts/whoami.exe.md`](./internal/posts/whoami.exe.md).
Any markdown file with `kind: "page"` follows the same routing rule automatically.
