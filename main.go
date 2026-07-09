package main

import (
	"log"
	"net/http"
	"os"

	"login-system/database"
	"login-system/handlers"
)

func main() {
	// Connect Database
	database.ConnectDB()

	// Routes
	http.HandleFunc("/", handlers.LoginPage)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/logout", handlers.Logout)

	// Static Files
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server Started")
	log.Println("http://localhost:" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
