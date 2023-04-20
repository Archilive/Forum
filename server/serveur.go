package main

import (
	"fmt"
	f "forum"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var url = "localhost:8080"

func main() {
	css, img, ico := http.FileServer(http.Dir("view/")), http.FileServer(http.Dir("Avatar/")), http.FileServer(http.Dir("assets/"))
	f.Init()
	http.Handle("/view/", http.StripPrefix("/view/", css))
	http.Handle("/Avatar/", http.StripPrefix("/Avatar/", img))
	http.Handle("/assets/", http.StripPrefix("/assets/", ico))
	http.HandleFunc("/", f.Index)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/logout", f.Logout)
	http.HandleFunc("/createPost", f.CreatePost)
	http.HandleFunc("/createComment", f.CreateComment)
	http.HandleFunc("/setting", f.Setting)
	http.HandleFunc("/register", f.Register)
	http.HandleFunc("/post/", f.Comments)
	http.HandleFunc("/like", f.LikeHandler)
	http.HandleFunc("/profile/", f.Profile)
	http.HandleFunc("/report", f.Report)
	http.HandleFunc("/requestMod", f.RequestMod)
	http.HandleFunc("/admin", f.Admin)
	http.HandleFunc("/edit/", f.Edit)
	http.HandleFunc("/createCategorie", f.CreateCategorie)
	http.HandleFunc("/acceptMod", f.AcceptMod)
	fmt.Println("Listening at http://" + url)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println(err)
	}
}
