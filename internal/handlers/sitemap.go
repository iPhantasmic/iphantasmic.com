package handlers

import (
	"encoding/xml"
	"net/http"

	"iphantasmic.com/internal/posts"
)

type sitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

func (s *Server) sitemap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urls := []sitemapURL{
		{Loc: s.site.URL("/")},
	}
	for _, post := range s.posts.SitemapPages() {
		urls = append(urls, sitemapURL{
			Loc:     s.pageURL(post),
			LastMod: post.Published.Format("2006-01-02"),
		})
	}

	writeXML(w, "application/xml; charset=utf-8", sitemapURLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	})
}

func (s *Server) pageURL(post posts.Post) string {
	if post.Kind == "page" {
		return s.site.URL("/" + post.Slug)
	}
	return s.site.URL("/posts/" + post.Slug)
}

func (s *Server) robots(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("User-agent: *\nAllow: /\n\nSitemap: " + s.site.URL("/sitemap.xml") + "\n"))
}

func writeXML(w http.ResponseWriter, contentType string, value any) {
	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	_ = encoder.Encode(value)
}
