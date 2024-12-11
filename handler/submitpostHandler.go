package handler

import (
	"encoding/json"
	"net/http"
)

// Gestionnaire API pour soumettre un post
// Fonction pour soumettre un post via API
func SubmitPostAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newPost Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		mu.Lock()
		posts = append(posts, newPost)
		mu.Unlock()

		// Répondre avec le post créé
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newPost)
		return
	}
	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
