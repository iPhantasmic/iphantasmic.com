package posts

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	urlpath "path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"gopkg.in/yaml.v3"
)

type Post struct {
	Title       string
	Slug        string
	Description string
	Kind        string
	Icon        string
	Cover       string
	Canonical   string
	Featured    bool
	Published   time.Time
	Updated     time.Time
	Tags        []string
	Body        template.HTML
	searchText  string
}

type Store struct {
	posts  []Post
	pages  []Post
	bySlug map[string]Post
}

type SearchResult struct {
	Post    Post
	Excerpt string
	Score   int
}

type frontMatter struct {
	Title       string   `yaml:"title"`
	Slug        string   `yaml:"slug"`
	Description string   `yaml:"description"`
	Kind        string   `yaml:"kind"`
	Icon        string   `yaml:"icon"`
	Cover       string   `yaml:"cover"`
	Canonical   string   `yaml:"canonical"`
	Featured    bool     `yaml:"featured"`
	Published   string   `yaml:"published"`
	Updated     string   `yaml:"updated"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
}

var (
	slugPattern       = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	schemePattern     = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*:`)
	contentBaseKey    = parser.NewContextKey()
	calloutMarkerKind = map[string]string{
		"[!NOTE]":      "note",
		"[!TIP]":       "tip",
		"[!IMPORTANT]": "important",
		"[!WARNING]":   "warning",
		"[!CAUTION]":   "caution",
	}
	calloutKindLabel = map[string]string{
		"note":      "Note",
		"tip":       "Tip",
		"important": "Important",
		"warning":   "Warning",
		"caution":   "Caution",
	}
	markdownRenderer = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(util.Prioritized(contentASTTransformer{}, 500)),
		),
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

	sort.SliceStable(loaded, func(i, j int) bool {
		if loaded[i].Featured != loaded[j].Featured {
			return loaded[i].Featured
		}
		return loaded[i].Published.After(loaded[j].Published)
	})

	bySlug := make(map[string]Post, len(loaded))
	posts := make([]Post, 0, len(loaded))
	pages := make([]Post, 0)
	for _, post := range loaded {
		if _, exists := bySlug[post.Slug]; exists {
			return nil, fmt.Errorf("duplicate post slug %q", post.Slug)
		}
		bySlug[post.Slug] = post
		if post.Kind == "page" {
			pages = append(pages, post)
		} else {
			posts = append(posts, post)
		}
	}

	return &Store{
		posts:  posts,
		pages:  pages,
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

func (s *Store) SitemapPages() []Post {
	items := make([]Post, 0, len(s.pages)+len(s.posts))
	items = append(items, s.pages...)
	items = append(items, s.posts...)
	return items
}

func (s *Store) Search(query string, limit int) []SearchResult {
	terms := searchTerms(query)
	if len(terms) == 0 {
		return nil
	}

	items := s.SitemapPages()
	results := make([]SearchResult, 0, len(items))
	for _, post := range items {
		score := post.searchScore(terms, query)
		if score == 0 {
			continue
		}

		results = append(results, SearchResult{
			Post:    post,
			Excerpt: excerpt(post.searchText, terms),
			Score:   score,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].Score != results[j].Score {
			return results[i].Score > results[j].Score
		}
		if results[i].Post.Featured != results[j].Post.Featured {
			return results[i].Post.Featured
		}
		return results[i].Post.Published.After(results[j].Post.Published)
	})

	if limit > 0 && len(results) > limit {
		return results[:limit]
	}
	return results
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

	slug := strings.TrimSpace(meta.Slug)
	contentBase := contentAssetBase(slug)
	updated, err := parseOptionalDate(path, "updated", meta.Updated)
	if err != nil {
		return Post{}, false, err
	}

	body = stripLeadingTitleHeading(body, meta.Title)
	searchText := plainMarkdownText(body)

	rendered, err := renderMarkdown(body, contentBase)
	if err != nil {
		return Post{}, false, fmt.Errorf("%s: render markdown: %w", path, err)
	}

	return Post{
		Title:       strings.TrimSpace(meta.Title),
		Slug:        slug,
		Description: strings.TrimSpace(meta.Description),
		Kind:        cleanKind(meta.Kind),
		Icon:        cleanIcon(meta.Icon, contentBase),
		Cover:       resolveContentAssetPath(contentBase, meta.Cover),
		Canonical:   strings.TrimSpace(meta.Canonical),
		Featured:    meta.Featured,
		Published:   published,
		Updated:     updated,
		Tags:        cleanTags(meta.Tags),
		Body:        rendered,
		searchText:  searchText,
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
		return fmt.Errorf("%s: slug %q must be lowercase kebab-case or dot-separated", path, meta.Slug)
	case strings.TrimSpace(meta.Description) == "":
		return fmt.Errorf("%s: missing required frontmatter field description", path)
	case strings.TrimSpace(meta.Published) == "":
		return fmt.Errorf("%s: missing required frontmatter field published", path)
	case cleanKind(meta.Kind) != "post" && cleanKind(meta.Kind) != "page":
		return fmt.Errorf("%s: kind must be post or page", path)
	}
	return nil
}

func cleanKind(kind string) string {
	kind = strings.TrimSpace(kind)
	if kind == "" {
		return "post"
	}
	return kind
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

func parseOptionalDate(path, field, value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, nil
	}

	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s: %s must use YYYY-MM-DD: %w", path, field, err)
	}
	return parsed, nil
}

func (p Post) LastModified() time.Time {
	if !p.Updated.IsZero() {
		return p.Updated
	}
	return p.Published
}

func (p Post) URLPath() string {
	if p.Kind == "page" {
		return "/" + p.Slug
	}
	return "/posts/" + p.Slug
}

func (p Post) searchScore(terms []string, query string) int {
	title := strings.ToLower(p.Title)
	description := strings.ToLower(p.Description)
	slug := strings.ToLower(p.Slug)
	body := strings.ToLower(p.searchText)
	tags := strings.ToLower(strings.Join(p.Tags, " "))
	exact := strings.ToLower(strings.TrimSpace(query))

	score := 0
	if exact != "" {
		score += weightedContains(title, exact, 90)
		score += weightedContains(tags, exact, 70)
		score += weightedContains(description, exact, 45)
		score += weightedContains(slug, exact, 35)
		score += weightedContains(body, exact, 20)
	}

	for _, term := range terms {
		score += weightedContains(title, term, 45)
		score += weightedContains(tags, term, 34)
		score += weightedContains(description, term, 22)
		score += weightedContains(slug, term, 16)
		score += weightedContains(body, term, 8)
	}

	if p.Featured && score > 0 {
		score += 3
	}
	return score
}

func weightedContains(value, term string, weight int) int {
	if value == "" || term == "" || !strings.Contains(value, term) {
		return 0
	}
	return weight
}

func renderMarkdown(raw []byte, contentBase string) (template.HTML, error) {
	var out bytes.Buffer

	context := parser.NewContext()
	context.Set(contentBaseKey, contentBase)
	reader := text.NewReader(raw)
	doc := markdownRenderer.Parser().Parse(reader, parser.WithContext(context))
	if err := markdownRenderer.Renderer().Render(&out, raw, doc); err != nil {
		return "", err
	}

	return template.HTML(stripRenderedCalloutMarkers(out.String())), nil
}

func plainMarkdownText(raw []byte) string {
	context := parser.NewContext()
	reader := text.NewReader(raw)
	doc := markdownRenderer.Parser().Parse(reader, parser.WithContext(context))

	var parts []string
	_ = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch typed := node.(type) {
		case *ast.Text:
			parts = append(parts, string(typed.Value(raw)))
		case *ast.String:
			parts = append(parts, string(typed.Value))
		case *ast.CodeBlock:
			parts = append(parts, string(typed.Text(raw)))
		case *ast.FencedCodeBlock:
			parts = append(parts, string(typed.Text(raw)))
		}
		return ast.WalkContinue, nil
	})

	return cleanSearchText(strings.Join(parts, " "))
}

func searchTerms(query string) []string {
	fields := strings.Fields(strings.ToLower(query))
	if len(fields) == 0 {
		return nil
	}

	terms := make([]string, 0, len(fields))
	seen := map[string]bool{}
	for _, field := range fields {
		field = strings.Trim(field, " \t\r\n.,;:!?()[]{}\"'`")
		if field == "" || seen[field] {
			continue
		}
		seen[field] = true
		terms = append(terms, field)
	}
	return terms
}

func cleanSearchText(value string) string {
	for marker := range calloutMarkerKind {
		value = strings.ReplaceAll(value, marker, " ")
	}
	return strings.Join(strings.Fields(value), " ")
}

func excerpt(value string, terms []string) string {
	const maxRunes = 180

	value = cleanSearchText(value)
	if value == "" {
		return ""
	}

	lower := strings.ToLower(value)
	start := 0
	for _, term := range terms {
		if index := strings.Index(lower, term); index >= 0 {
			start = index - 52
			if start < 0 {
				start = 0
			}
			break
		}
	}

	prefix := ""
	if start > 0 {
		if nextSpace := strings.IndexByte(value[start:], ' '); nextSpace >= 0 {
			start += nextSpace + 1
		}
		prefix = "..."
	}

	runes := []rune(value[start:])
	if len(runes) <= maxRunes {
		return prefix + string(runes)
	}

	return prefix + string(runes[:maxRunes]) + "..."
}

type contentASTTransformer struct{}

func (contentASTTransformer) Transform(node *ast.Document, reader text.Reader, context parser.Context) {
	contentBase, _ := context.Get(contentBaseKey).(string)
	source := reader.Source()

	_ = ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch typed := node.(type) {
		case *ast.Image:
			typed.Destination = []byte(resolveContentAssetPath(contentBase, string(typed.Destination)))
		case *ast.Blockquote:
			classifyCallout(typed, source)
		}

		return ast.WalkContinue, nil
	})
}

func classifyCallout(blockquote *ast.Blockquote, source []byte) {
	paragraph, ok := blockquote.FirstChild().(*ast.Paragraph)
	if !ok {
		return
	}

	text := strings.TrimSpace(string(paragraph.Text(source)))
	for marker, kind := range calloutMarkerKind {
		if !strings.HasPrefix(text, marker) {
			continue
		}

		blockquote.SetAttributeString("class", "callout callout-"+kind)
		blockquote.SetAttributeString("data-callout", calloutKindLabel[kind])
		return
	}
}

func contentAssetBase(slug string) string {
	return "/static/content/" + slug
}

func stripRenderedCalloutMarkers(value string) string {
	for marker := range calloutMarkerKind {
		value = strings.ReplaceAll(value, "<p>"+marker+"\n", "<p>")
		value = strings.ReplaceAll(value, "<p>"+marker+"<br />\n", "<p>")
		value = strings.ReplaceAll(value, "<p>"+marker+"</p>\n", "")
	}
	return value
}

func cleanIcon(icon, contentBase string) string {
	icon = strings.TrimSpace(icon)
	if !looksLikeImagePath(icon) {
		return icon
	}
	return resolveContentAssetPath(contentBase, icon)
}

func looksLikeImagePath(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	if strings.HasPrefix(value, "/") || schemePattern.MatchString(value) {
		return true
	}

	assetPath, _, _ := strings.Cut(value, "?")
	assetPath, _, _ = strings.Cut(assetPath, "#")
	switch strings.ToLower(urlpath.Ext(assetPath)) {
	case ".avif", ".gif", ".ico", ".jpeg", ".jpg", ".png", ".svg", ".webp":
		return true
	default:
		return false
	}
}

func resolveContentAssetPath(contentBase, value string) string {
	value = strings.TrimSpace(value)
	if value == "" || contentBase == "" || strings.HasPrefix(value, "/") || strings.HasPrefix(value, "#") || schemePattern.MatchString(value) {
		return value
	}

	pathPart := value
	suffix := ""
	if index := strings.IndexAny(value, "?#"); index >= 0 {
		pathPart = value[:index]
		suffix = value[index:]
	}

	clean := urlpath.Clean("/" + strings.TrimPrefix(pathPart, "./"))
	if clean == "/" {
		return value
	}
	return contentBase + clean + suffix
}

func stripLeadingTitleHeading(raw []byte, title string) []byte {
	title = normalizeHeadingText(title)
	if title == "" {
		return raw
	}

	body := bytes.ReplaceAll(raw, []byte("\r\n"), []byte("\n"))
	body = bytes.TrimLeft(body, "\n\t ")
	lines := bytes.SplitN(body, []byte("\n"), 2)
	first := strings.TrimSpace(string(lines[0]))
	if !strings.HasPrefix(first, "# ") {
		return raw
	}

	heading := normalizeHeadingText(strings.TrimSpace(strings.TrimPrefix(first, "# ")))
	if heading != title {
		return raw
	}

	if len(lines) == 1 {
		return nil
	}

	return bytes.TrimLeft(lines[1], "\n\t ")
}

func normalizeHeadingText(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "*_`")
	return strings.TrimSpace(value)
}
