package handler

import (
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create_post.html", nil)
}
