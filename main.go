package main

import (
	"fmt"
	"log"
	"net/http"

	"Forum/handler"
)

func main() {
	darkBlue := "\033[38;5;17m" // Dark Blue (#0A192F)
	cyan := "\033[38;5;39m"     // Cyan (#64ffda)
	reset := "\033[0m"
	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/forum", handler.ForumHandler)
	http.HandleFunc("/create_post", handler.CreatePostHandler)
	http.HandleFunc("/submitpost", handler.SubmitPostAPI)
	fmt.Println(darkBlue + "/-------------------------------------------------------------------\\" + reset)
	fmt.Println(darkBlue + "|" + cyan + "                Server started on http://localhost:8080            " + darkBlue + " |" + reset)
	fmt.Println(darkBlue + "\\-------------------------------------------------------------------/" + reset)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
