package main

import (
	"fmt"
	fh "forum/Handler"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", fh.HandleIndex)
	http.HandleFunc("/register", fh.HandleRegister)
	http.HandleFunc("/login", fh.HandleLogin)
	http.HandleFunc("/delete", fh.HandleDeleteTopic)
	http.HandleFunc("/modifyt", fh.HandleModifyTopic)
	http.HandleFunc("/create", fh.HandleCreate)
	http.HandleFunc("/profile", fh.HandleProfile)
	http.HandleFunc("/logout", fh.HandleLogout)
	http.HandleFunc("/infos", fh.HandleInfos)
	http.HandleFunc("/test", fh.HandleTest)
	http.HandleFunc("/editprofile", fh.HandleEditProfile)
	http.HandleFunc("/addcomment", fh.HandleAddComment)
	http.HandleFunc("/delcomment", fh.HandleDeleteComment)
	http.HandleFunc("/modifc", fh.HandleModifyComment)
	http.HandleFunc("/like", fh.HandleLike)
	http.HandleFunc("/dislike", fh.HandleDislike)
	http.HandleFunc("/notif", fh.HandleNotif)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	return
}
