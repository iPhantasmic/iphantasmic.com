package posts

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Post struct {
	Title       string
	Slug        string
	Description string
	Published   time.Time
	Tags        []string
	Body        template.HTML
}

type Store struct {
	posts  []Post
	bySlug map[string]Post
}

func LoadDir(dir string) (*Store, error) {
	var loaded []Post

	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		post, err := loadPost(path)
		if err != nil {
			return err
		}
		loaded = append(loaded, post)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(loaded, func(i, j int) bool {
		return loaded[i].Published.After(loaded[j].Published)
	})

	bySlug := make(map[string]Post, len(loaded))
	for _, post := range loaded {
		bySlug[post.Slug] = post
	}

	return &Store{
		posts:  loaded,
		bySlug: bySlug,
	}, nil
}

func (s *Store) All() []Post {
	return append([]Post(nil), s.posts...)
}

func (s *Store) Find(slug string) (Post, bool) {
	post, ok := s.bySlug[slug]
	return post, ok
}

func loadPost(path string) (Post, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}

	meta, body := splitFrontMatter(string(raw))
	slug := meta["slug"]
	if slug == "" {
		slug = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}

	published := time.Now()
	if value := meta["published"]; value != "" {
		parsed, err := time.Parse("2006-01-02", value)
		if err != nil {
			return Post{}, fmt.Errorf("%s: parse published: %w", path, err)
		}
		published = parsed
	}

	title := meta["title"]
	if title == "" {
		title = strings.ReplaceAll(slug, "-", " ")
	}

	return Post{
		Title:       title,
		Slug:        slug,
		Description: meta["description"],
		Published:   published,
		Tags:        parseTags(meta["tags"]),
		Body:        renderMarkdown(body),
	}, nil
}

func splitFrontMatter(raw string) (map[string]string, string) {
	meta := map[string]string{}
	raw = strings.ReplaceAll(raw, "\r\n", "\n")

	if !strings.HasPrefix(raw, "---\n") {
		return meta, raw
	}

	rest := strings.TrimPrefix(raw, "---\n")
	parts := strings.SplitN(rest, "\n---\n", 2)
	if len(parts) != 2 {
		return meta, raw
	}

	for _, line := range strings.Split(parts[0], "\n") {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), "\"")
		meta[key] = value
	}

	return meta, parts[1]
}

func parseTags(raw string) []string {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "[")
	raw = strings.TrimSuffix(raw, "]")
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	tags := make([]string, 0, len(parts))
	for _, part := range parts {
		tag := strings.Trim(strings.TrimSpace(part), "\"")
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func renderMarkdown(raw string) template.HTML {
	lines := strings.Split(strings.ReplaceAll(raw, "\r\n", "\n"), "\n")
	var out bytes.Buffer
	var paragraph []string
	inCode := false
	inList := false

	flushParagraph := func() {
		if len(paragraph) == 0 {
			return
		}
		out.WriteString("<p>")
		out.WriteString(renderInline(strings.Join(paragraph, " ")))
		out.WriteString("</p>\n")
		paragraph = nil
	}

	flushList := func() {
		if inList {
			out.WriteString("</ul>\n")
			inList = false
		}
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "```") {
			flushParagraph()
			flushList()
			if inCode {
				out.WriteString("</code></pre>\n")
				inCode = false
			} else {
				out.WriteString("<pre><code>")
				inCode = true
			}
			continue
		}

		if inCode {
			out.WriteString(html.EscapeString(line))
			out.WriteByte('\n')
			continue
		}

		if trimmed == "" {
			flushParagraph()
			flushList()
			continue
		}

		switch {
		case strings.HasPrefix(trimmed, "### "):
			flushParagraph()
			flushList()
			out.WriteString("<h3>")
			out.WriteString(renderInline(strings.TrimPrefix(trimmed, "### ")))
			out.WriteString("</h3>\n")
		case strings.HasPrefix(trimmed, "## "):
			flushParagraph()
			flushList()
			out.WriteString("<h2>")
			out.WriteString(renderInline(strings.TrimPrefix(trimmed, "## ")))
			out.WriteString("</h2>\n")
		case strings.HasPrefix(trimmed, "# "):
			flushParagraph()
			flushList()
			out.WriteString("<h1>")
			out.WriteString(renderInline(strings.TrimPrefix(trimmed, "# ")))
			out.WriteString("</h1>\n")
		case strings.HasPrefix(trimmed, "> "):
			flushParagraph()
			flushList()
			out.WriteString("<blockquote>")
			out.WriteString(renderInline(strings.TrimPrefix(trimmed, "> ")))
			out.WriteString("</blockquote>\n")
		case strings.HasPrefix(trimmed, "- "):
			flushParagraph()
			if !inList {
				out.WriteString("<ul>\n")
				inList = true
			}
			out.WriteString("<li>")
			out.WriteString(renderInline(strings.TrimPrefix(trimmed, "- ")))
			out.WriteString("</li>\n")
		case trimmed == "---":
			flushParagraph()
			flushList()
			out.WriteString("<hr>\n")
		default:
			flushList()
			paragraph = append(paragraph, trimmed)
		}
	}

	flushParagraph()
	flushList()
	if inCode {
		out.WriteString("</code></pre>\n")
	}

	return template.HTML(out.String())
}

var (
	imageRe      = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	linkRe       = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	boldRe       = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	inlineCodeRe = regexp.MustCompile("`([^`]+)`")
)

func renderInline(raw string) string {
	escaped := html.EscapeString(raw)
	escaped = imageRe.ReplaceAllString(escaped, `<img src="$2" alt="$1">`)
	escaped = linkRe.ReplaceAllString(escaped, `<a href="$2">$1</a>`)
	escaped = boldRe.ReplaceAllString(escaped, `<strong>$1</strong>`)
	escaped = inlineCodeRe.ReplaceAllString(escaped, `<code>$1</code>`)
	return escaped
}
