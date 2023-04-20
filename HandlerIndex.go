package forum

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var t = template.Must(template.ParseGlob("view/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	var template TemplatePost
	fm, _ := GetFlash(w, r, "message")
	template.Categories = GetALLCategories()
	if CheckCookie(r) {
		template.User = GetUser(r)
		template.Posts = GetAllPost(template.User.Id)
		template.IsLogged = true
	} else {
		template.IsLogged = false
		template.Posts = GetAllPost(-1)
	}
	if r.FormValue("Categorie") != "All" {
		for _, cat := range template.Categories {
			if cat.Name == r.FormValue("Categorie") {
				template.Categorie = cat.Name
				var newtable []Post
				for _, post := range template.Posts {
					if IntInSlice(cat.Id, post.Categories) {
						newtable = append(newtable, post)
					}
				}
				template.Posts = newtable
			}
		}
	} else {
		template.Categorie = "All"
	}
	switch r.FormValue("Filter") {
	case "Like":
		template.Filter = "Like"
		if r.FormValue("Type") == "Ascending" {
			template.Order = "Ascending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].LikeCount < template.Posts[j].LikeCount
			})
		} else if r.FormValue("Type") == "Descending" {
			template.Order = "Descending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].LikeCount > template.Posts[j].LikeCount
			})
		}
	case "DisLike":
		template.Filter = "DisLike"
		if r.FormValue("Type") == "Ascending" {
			template.Order = "Ascending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].DisLikeCount < template.Posts[j].DisLikeCount
			})
		} else if r.FormValue("Type") == "Descending" {
			template.Order = "Descending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].DisLikeCount > template.Posts[j].DisLikeCount
			})
		}
	case "Comments":
		template.Filter = "Comments"
		if r.FormValue("Type") == "Ascending" {
			template.Order = "Ascending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].NbrComments < template.Posts[j].NbrComments
			})
		} else if r.FormValue("Type") == "Descending" {
			template.Order = "Descending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].NbrComments > template.Posts[j].NbrComments
			})
		}
	case "Date":
		template.Filter = "Date"
		if r.FormValue("Type") == "Ascending" {
			template.Order = "Ascending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].UpdatedAt.Before(template.Posts[j].UpdatedAt)
			})
		} else {
			template.Order = "Descending"
			sort.Slice(template.Posts, func(i, j int) bool {
				return template.Posts[i].UpdatedAt.After(template.Posts[j].UpdatedAt)
			})
		}
	}
	t.ExecuteTemplate(w, "index.html", template)
	fmt.Fprintf(w, "%s", fm)
}

func Comments(w http.ResponseWriter, r *http.Request) {
	post_id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/post/"))
	var template TemplateComment
	fm, _ := GetFlash(w, r, "message")
	if CheckCookie(r) {
		template.User = GetUser(r)
		template.IsLogged = true
		template.Comments = GetAllComment(post_id, template.User.Id)
		template.Post = GetPost(post_id, template.User.Id)
	} else {
		template.IsLogged = false
		template.Comments = GetAllComment(post_id, -1)
		template.Post = GetPost(post_id, -1)
	}
	t.ExecuteTemplate(w, "post.html", template)
	fmt.Fprintf(w, "%s", fm)
}
