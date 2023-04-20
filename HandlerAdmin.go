package forum

import (
	"fmt"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		user := GetUser(r)
		if user.IsAdmin {
			switch r.Method {
			case "GET":
				var template TemplateAdmin
				template.IsLogged = true
				template.User = user
				template.ReportInfo = GetReport()
				template.RequestMod = GetRequest()
				fm, _ := GetFlash(w, r, "message")
				t.ExecuteTemplate(w, "admin.html", template)
				fmt.Fprintf(w, "%s", fm)
			case "POST":
				id, _ := strconv.Atoi(r.FormValue("id"))
				switch r.FormValue("type") {
				case "Post":
					switch r.FormValue("action") {
					case "SUPPRIMER":
						DeletePost(id)
						SetFlash(w, "message", []byte("Post deleted"))
					case "ANNULER":
						CancelReportPost(id)
						SetFlash(w, "message", []byte("Report Cancelled"))
					}
				case "Comment":
					switch r.FormValue("action") {
					case "SUPPRIMER":
						DeleteComment(id)
						SetFlash(w, "message", []byte("Comment deleted"))
					case "ANNULER":
						CancelReportComment(id)
						SetFlash(w, "message", []byte("Report Cancelled"))
					}
				}
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}
		} else {
			SetFlash(w, "message", []byte("You are not able to go here"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Report(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		SetFlash(w, "message", []byte("Something went wrong"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		if CheckCookie(r) {
			user := GetUser(r)
			if user.IsMod {
				table := r.FormValue("table")
				id, _ := strconv.Atoi(r.FormValue("id"))
				switch table {
				case "ReportPost":
					AddReportPost(id, user.Id)
					SetFlash(w, "message", []byte("Post reported"))
					http.Redirect(w, r, "/", http.StatusSeeOther)
				case "ReportComment":
					post_id := GetPostId(id)
					AddReportComment(post_id, id, user.Id)
					SetFlash(w, "message", []byte("Comment reported"))
					http.Redirect(w, r, "/post/"+strconv.Itoa(post_id), http.StatusSeeOther)
				}
			} else {
				SetFlash(w, "message", []byte("You are not a Mod"))
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func RequestMod(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		user := GetUser(r)
		var template TemplateCreatePost
		template.User = user
		template.IsLogged = true
		if !template.User.IsMod {
			switch r.Method {
			case "GET":
				fm, _ := GetFlash(w, r, "message")
				t.ExecuteTemplate(w, "mod.html", template)
				fmt.Fprintf(w, "%s", fm)
			case "POST":
				Reason := r.FormValue("Reason")
				if AddRequestMod(template.User.Id, Reason) {
					SetFlash(w, "message", []byte("Request to become mod posted"))
					http.Redirect(w, r, "/", http.StatusSeeOther)
				} else {
					SetFlash(w, "message", []byte("You already requested to become mod"))
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		} else {
			SetFlash(w, "message", []byte("You are already a Mod"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func AcceptMod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		SetFlash(w, "message", []byte("Something went wrong"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		if CheckCookie(r) {
			user := GetUser(r)
			if user.IsAdmin {
				switch r.FormValue("action") {
				case "accept":
					RankupMod(r.FormValue("username"))
					fmt.Println(r.FormValue("username"))
					DeleteRequest(r.FormValue("username"))
				case "denied":
					DeleteRequest(r.FormValue("username"))
				}
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			} else {
				SetFlash(w, "message", []byte("You are not admin"))
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
