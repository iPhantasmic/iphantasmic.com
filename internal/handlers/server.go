package handlers

import (
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/a-h/templ"

	"iphantasmic.com/internal/config"
	"iphantasmic.com/internal/posts"
	"iphantasmic.com/internal/templates"
)

type Server struct {
	site  config.Site
	posts *posts.Store
}

func NewServer(site config.Site, posts *posts.Store) *Server {
	return &Server{
		site:  site,
		posts: posts,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/favicon.ico", s.favicon)
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/feed.xml", s.feed)
	mux.HandleFunc("/sitemap.xml", s.sitemap)
	mux.HandleFunc("/robots.txt", s.robots)
	mux.HandleFunc("/search", s.search)
	mux.HandleFunc("/tags/", s.tag)
	mux.HandleFunc("/tags", s.tags)
	mux.HandleFunc("/posts/", s.post)
	mux.HandleFunc("/", s.root)
	return mux
}

func (s *Server) root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		s.render(w, r, http.StatusOK, templates.Index(s.site, time.Now().Year(), s.posts.All()))
		return
	}

	slug := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	page, ok := s.posts.Find(slug)
	if !ok || page.Kind != "page" {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.PostPage(s.site, time.Now().Year(), page))
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(path.Clean(r.URL.Path), "/posts/")
	if slug == "" || slug == "." {
		s.notFound(w, r)
		return
	}

	post, ok := s.posts.Find(slug)
	if !ok || post.Kind != "post" {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.PostPage(s.site, time.Now().Year(), post))
}

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	results := s.posts.Search(query, 20)
	if r.URL.Query().Get("partial") == "1" {
		s.render(w, r, http.StatusOK, templates.SearchResults(query, results))
		return
	}

	s.render(w, r, http.StatusOK, templates.Search(s.site, time.Now().Year(), query, results))
}

func (s *Server) tags(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/tags" {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.Tags(s.site, time.Now().Year(), s.posts.Tags()))
}

func (s *Server) tag(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := strings.TrimPrefix(path.Clean(r.URL.Path), "/tags/")
	if slug == "" || slug == "." {
		http.Redirect(w, r, "/tags", http.StatusMovedPermanently)
		return
	}

	tag, taggedPosts, ok := s.posts.PostsByTag(slug)
	if !ok {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.TagPage(s.site, time.Now().Year(), tag, taggedPosts))
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusNotFound, templates.NotFound(s.site, time.Now().Year()))
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func (s *Server) favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func (s *Server) render(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
