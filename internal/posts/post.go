package posts

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
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

type frontMatter struct {
	Title       string   `yaml:"title"`
	Slug        string   `yaml:"slug"`
	Description string   `yaml:"description"`
	Published   string   `yaml:"published"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
}

var (
	slugPattern      = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	markdownRenderer = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(goldmarkhtml.WithXHTML()),
	)
)

func LoadDir(dir string) (*Store, error) {
	var loaded []Post

	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		post, include, err := loadPost(path)
		if err != nil {
			return err
		}
		if include {
			loaded = append(loaded, post)
		}
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
		if _, exists := bySlug[post.Slug]; exists {
			return nil, fmt.Errorf("duplicate post slug %q", post.Slug)
		}
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

func loadPost(path string) (Post, bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Post{}, false, err
	}

	metaBytes, body, err := splitFrontMatter(raw)
	if err != nil {
		return Post{}, false, fmt.Errorf("%s: %w", path, err)
	}

	var meta frontMatter
	if err := yaml.Unmarshal(metaBytes, &meta); err != nil {
		return Post{}, false, fmt.Errorf("%s: parse frontmatter: %w", path, err)
	}
	if meta.Draft {
		return Post{}, false, nil
	}

	if err := validateFrontMatter(path, meta); err != nil {
		return Post{}, false, err
	}

	published, err := time.Parse("2006-01-02", meta.Published)
	if err != nil {
		return Post{}, false, fmt.Errorf("%s: published must use YYYY-MM-DD: %w", path, err)
	}

	rendered, err := renderMarkdown(body)
	if err != nil {
		return Post{}, false, fmt.Errorf("%s: render markdown: %w", path, err)
	}

	return Post{
		Title:       strings.TrimSpace(meta.Title),
		Slug:        strings.TrimSpace(meta.Slug),
		Description: strings.TrimSpace(meta.Description),
		Published:   published,
		Tags:        cleanTags(meta.Tags),
		Body:        rendered,
	}, true, nil
}

func splitFrontMatter(raw []byte) ([]byte, []byte, error) {
	raw = bytes.ReplaceAll(raw, []byte("\r\n"), []byte("\n"))
	if !bytes.HasPrefix(raw, []byte("---\n")) {
		return nil, nil, fmt.Errorf("missing YAML frontmatter")
	}

	rest := bytes.TrimPrefix(raw, []byte("---\n"))
	parts := bytes.SplitN(rest, []byte("\n---\n"), 2)
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("unterminated YAML frontmatter")
	}

	return parts[0], parts[1], nil
}

func validateFrontMatter(path string, meta frontMatter) error {
	switch {
	case strings.TrimSpace(meta.Title) == "":
		return fmt.Errorf("%s: missing required frontmatter field title", path)
	case strings.TrimSpace(meta.Slug) == "":
		return fmt.Errorf("%s: missing required frontmatter field slug", path)
	case !slugPattern.MatchString(strings.TrimSpace(meta.Slug)):
		return fmt.Errorf("%s: slug %q must be lowercase kebab-case", path, meta.Slug)
	case strings.TrimSpace(meta.Description) == "":
		return fmt.Errorf("%s: missing required frontmatter field description", path)
	case strings.TrimSpace(meta.Published) == "":
		return fmt.Errorf("%s: missing required frontmatter field published", path)
	}
	return nil
}

func cleanTags(tags []string) []string {
	if len(tags) == 0 {
		return nil
	}

	cleaned := make([]string, 0, len(tags))
	seen := map[string]bool{}
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" || seen[tag] {
			continue
		}
		seen[tag] = true
		cleaned = append(cleaned, tag)
	}
	return cleaned
}

func renderMarkdown(raw []byte) (template.HTML, error) {
	var out bytes.Buffer
	if err := markdownRenderer.Convert(raw, &out); err != nil {
		return "", err
	}

	return template.HTML(out.String()), nil
}
