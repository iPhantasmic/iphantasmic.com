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
	mux.HandleFunc("/whoami.exe", s.page)
	mux.HandleFunc("/posts/", s.post)
	mux.HandleFunc("/", s.index)
	return mux
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.Index(s.site, time.Now().Year(), s.posts.All()))
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

func (s *Server) page(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	page, ok := s.posts.Find(slug)
	if !ok || page.Kind != "page" {
		s.notFound(w, r)
		return
	}

	s.render(w, r, http.StatusOK, templates.PostPage(s.site, time.Now().Year(), page))
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
