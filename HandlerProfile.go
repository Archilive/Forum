package forum

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	user_id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/profile/"))
	var template TemplateProfile
	template.UserInfo = GetAllUserInfo(user_id)
	fm, _ := GetFlash(w, r, "message")
	if CheckCookie(r) {
		template.User = GetUser(r)
		template.IsLogged = true
		t.ExecuteTemplate(w, "profile.html", template)
		fmt.Fprintf(w, "%s", fm)
	} else {
		template.IsLogged = false
		t.ExecuteTemplate(w, "profile.html", template)
		fmt.Fprintf(w, "%s", fm)
	}
}

func Setting(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		fm, _ := GetFlash(w, r, "message")
		var template TemplateSetting
		template.User = GetUser(r)
		template.IsLogged = true
		switch r.Method {
		case "GET":
			t.ExecuteTemplate(w, "setting.html", template)
			fmt.Fprintf(w, "%s", fm)
		case "POST":
			if len(r.FormValue("Username")) > 0 {
				_, err := db.Exec("UPDATE User Set Username = ? WHERE Id = ?", r.FormValue("Username"), template.User.Id)
				if err != nil {
					SetFlash(w, "message", []byte("Username already used"))
				} else {
					template.User.Username = r.FormValue("Username")
				}
			}
			if len(r.FormValue("Password")) > 0 {
				template.User.Password, _ = HashPassword(r.FormValue("Password"))
				db.Exec("UPDATE User Set Password = ? WHERE Id = ?", template.User.Password, template.User.Id)
			}
			if len(r.FormValue("Email")) > 0 {
				_, err := db.Exec("UPDATE User Set Email = ? WHERE Id = ?", r.FormValue("Email"), template.User.Id)
				if err != nil {
					SetFlash(w, "message", []byte("Password already used"))
				} else {
					template.User.Email = r.FormValue("Email")
				}
			}
			Uploadfile(w, r, template.User.Id)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Fprintf(w, "%s", fm)
		}
	} else {
		SetFlash(w, "message", []byte("You need to be logged in to go to setting"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Uploadfile(w http.ResponseWriter, r *http.Request, User_id int) {
	err1 := r.ParseMultipartForm(10 << 20)
	if err1 != nil {
		SetFlash(w, "message", []byte("Image is too heavy"))
		http.Redirect(w, r, "/setting", http.StatusSeeOther)
		return
	}
	file, _, _ := r.FormFile("MyAvatar")
	if file != nil {
		defer file.Close()
		StrUser_id := strconv.Itoa(User_id)
		os.Remove("Avatar/" + StrUser_id + ".png")
		dst, _ := os.Create("Avatar/" + StrUser_id + ".png")
		io.Copy(dst, file)
		defer dst.Close()
	}
}
