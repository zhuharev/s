package shortener

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"strings"
)

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.redirectShortURL(w, r)
	case "POST":
		s.createShortURL(w, r)
	}
}

func (s *Service) redirectShortURL(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, s.Config.Web.Endpoint)
	url, err := s.Store.Get(code)
	if err != nil || url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, 301)
}

func (s *Service) createShortURL(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL must be provided", 400)
		return
	}
	customCode := r.FormValue("code")
	code, e := s.CreateShortUrl(url, customCode)
	if e != nil {
		http.Error(w, "Some error", 500)
	}

	fmt.Fprintf(w, "http://localhost:8080/%s\n", string(code))
}

func (s *Service) CreateShortUrl(u string, codes ...string) (string, error) {
	code := randStr(4)
	if len(codes) != 0 {
		code = codes[0]
	}
	for ok, e := s.Store.Has(code); code == "" || (ok || e != nil); code = randStr(4) {
	}
	e := s.Store.Set(code, u)
	color.Green("Created short url %s for %s", code, u)
	if e != nil {
		return "", e
	}
	return code, nil
}
