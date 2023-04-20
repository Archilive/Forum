package forum

import (
	"net/http"
	"strconv"
	"strings"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		user := GetUser(r)
		post_id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/edit/"))
		if IsOwner(post_id, user.Id) {
			switch r.Method {
			case "GET":
				var template TemplateEdit
				template.IsLogged = true
				template.User = user
				template.Post = GetPost(post_id, user.Id)
				Categories := GetALLCategories()
				for _, categorie := range Categories {
					if !(IntInSlice(categorie.Id, template.Post.Categories)) {
						template.Categories = append(template.Categories, categorie)
					}
				}
				t.ExecuteTemplate(w, "edit.html", template)
			case "POST":
				switch r.FormValue("action") {
				case "edit":
					EditPost(post_id, r.FormValue("Title"), r.FormValue("Text"))
					DeleteCategories(post_id)
					if err := r.ParseForm(); err != nil {
						return
					}
					for index, valeur := range r.PostForm {
						if index == "Categorie" {
							AddCategories(post_id, valeur)
						}
					}
					http.Redirect(w, r, "/post/"+strconv.Itoa(post_id), http.StatusSeeOther)
				case "delete":
					DeletePost(post_id)
					SetFlash(w, "message", []byte("Post successfully deleted"))
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		} else {
			SetFlash(w, "message", []byte("You are not the author of this post"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		SetFlash(w, "message", []byte("You need to be logged in to access to edit menu"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
