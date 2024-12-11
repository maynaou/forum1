package handler

import (
	"net/http"
	"sync"
)

type Post struct {
	UserID      int    `json:"user_id"`
	PostCreator string `json:"post_creator"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}

// Variable globale pour stocker les posts
var (
	posts []Post
	mu    sync.Mutex // Mutex pour protéger l'accès concurrent à posts
)

// Fonction pour afficher la page du forum
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Afficher la page du forum avec les posts
	templates.ExecuteTemplate(w, "forum.html", posts)
}
