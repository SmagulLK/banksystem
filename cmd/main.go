package main

import (
	"bank/app"
	"bank/db"
	"bank/postgres"
	"log"
	"net/http"
)

func main() {
	conn, err := db.NewConnection("postgres", "Admin", "Turing")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	repo := postgres.New(conn)
	application := app.NewApplication(repo)
	http.HandleFunc("/user", application.GetUserByName)
	http.HandleFunc("/", application.CreateUser)
	http.ListenAndServe(":7070", nil)
}
