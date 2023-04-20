package forum

import (
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		var template TemplateCreatePost
		template.User = GetUser(r)
		template.Categories = GetALLCategories()
		template.IsLogged = true
		switch r.Method {
		case "GET":
			t.ExecuteTemplate(w, "createPost.html", template)
		case "POST":
			Post_id := AddPost(template.User.Id, r.FormValue("Title"), r.FormValue("Text"))
			if Post_id == -1 {
				http.Redirect(w, r, "/createPost", http.StatusSeeOther)
			} else {
				if err := r.ParseForm(); err != nil {
					return
				}
				for index, valeur := range r.PostForm {
					if index == "Categorie" {
						AddCategories(Post_id, valeur)
					}
				}
				http.Redirect(w, r, "/post/"+strconv.Itoa(Post_id), http.StatusSeeOther)
			}
		}
	} else {
		SetFlash(w, "message", []byte("You need to be logged in to create a post"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		strPost_id := r.FormValue("Id")
		if CheckCookie(r) {
			user := GetUser(r)
			Post_id, _ := strconv.Atoi(strPost_id)
			AddComment(Post_id, user.Id, r.FormValue("Text"))
			http.Redirect(w, r, "/post/"+strPost_id, http.StatusSeeOther)
		} else {
			SetFlash(w, "message", []byte("You need to be logged in to create a comment"))
			http.Redirect(w, r, "/post/"+strPost_id, http.StatusSeeOther)
		}
	}
}

func CreateCategorie(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		user := GetUser(r)
		if user.IsAdmin {
			switch r.Method {
			case "GET":
				var template TemplateCreatePost
				template.User = user
				template.IsLogged = true
				template.Categories = GetALLCategories()
				t.ExecuteTemplate(w, "createCategorie.html", template)
			case "POST":
				AddCategorie(r.FormValue("Title"))
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
