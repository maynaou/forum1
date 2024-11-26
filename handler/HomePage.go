package handler

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ShowErrorPage(w, "Page not found", http.StatusNotFound)
		return
	}

	err := templates.ExecuteTemplate(w, "homePage.html", nil)
	if err != nil {
		ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
