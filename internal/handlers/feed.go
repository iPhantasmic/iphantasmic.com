package handlers

import (
	"encoding/xml"
	"net/http"
	"time"
)

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	LastBuildDate string    `xml:"lastBuildDate,omitempty"`
	Items         []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func (s *Server) feed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts := s.posts.All()
	items := make([]rssItem, 0, len(posts))
	lastModified := time.Time{}
	for _, post := range posts {
		link := s.site.URL(post.URLPath())
		items = append(items, rssItem{
			Title:       post.Title,
			Link:        link,
			GUID:        link,
			PubDate:     post.Published.Format(time.RFC1123Z),
			Description: post.Description,
		})
		if post.LastModified().After(lastModified) {
			lastModified = post.LastModified()
		}
	}

	lastBuildDate := time.Now().Format(time.RFC1123Z)
	if !lastModified.IsZero() {
		lastBuildDate = lastModified.Format(time.RFC1123Z)
	}

	writeXML(w, "application/rss+xml; charset=utf-8", rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         s.site.Name,
			Link:          s.site.URL("/"),
			Description:   s.site.Description,
			Language:      "en",
			LastBuildDate: lastBuildDate,
			Items:         items,
		},
	})
}
