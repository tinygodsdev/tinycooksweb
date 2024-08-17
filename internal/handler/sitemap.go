package handler

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

type SitemapURL struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	XmlNS   string       `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

func (h *Handler) generateSitemap(ctx context.Context) ([]byte, error) {
	lang := locale.Default()
	urls := []SitemapURL{
		{Loc: h.app.Cfg.BaseURL + "/", ChangeFreq: "weekly", Priority: "1.0"},
		{Loc: h.app.Cfg.BaseURL + "/about", ChangeFreq: "monthly", Priority: "0.5"},
		{Loc: h.app.Cfg.BaseURL + "/catalog", ChangeFreq: "weekly", Priority: "0.8"},
	}

	recipes, err := h.app.GetRecipes(ctx, recipe.Filter{Locale: lang})
	if err != nil {
		return nil, err
	}

	for _, recipe := range recipes {
		urls = append(urls, SitemapURL{
			Loc:      h.app.Cfg.BaseURL + "/recipe/" + recipe.Slug,
			Priority: "0.8",
		})
	}

	tags, err := h.app.GetTags(ctx, lang)
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		urls = append(urls, SitemapURL{
			Loc:      h.app.Cfg.BaseURL + "/tag/" + tag.Slug,
			Priority: "0.6",
		})
	}

	sitemap := Sitemap{
		XmlNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	return xml.MarshalIndent(sitemap, "", "  ")
}

func (h *Handler) SitemapHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
		xmlData, err := h.generateSitemap(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(xmlData)
	}
}

func (h *Handler) RobotsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		robotsContent := `User-agent: *
Disallow: /terms
Disallow: /privacy
Sitemap: ` + h.app.Cfg.BaseURL + `/sitemap.xml`

		w.Write([]byte(robotsContent))
	}
}
