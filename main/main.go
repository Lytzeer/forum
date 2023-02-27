package main

import (
	"database/sql"
	"fmt"
	fh "forum/Handler"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", fh.HandleIndex)
	http.HandleFunc("/register", fh.HandleRegister)
	http.HandleFunc("/login", fh.HandleLogin)
	http.HandleFunc("/create", fh.HandleCreate)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	sql.Open("sqlite3", "./database/forum.db")
	return
}
