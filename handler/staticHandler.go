package handler

import (
	"net/http"
	"os"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	_, err := os.Stat("." + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		ShowErrorPage(w, "don't have access", http.StatusForbidden)
		return
	}
	fs.ServeHTTP(w, r)
}
