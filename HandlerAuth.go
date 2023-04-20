package forum

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	uuid "github.com/google/uuid"
)

type Login_struct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register_struct struct {
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Image    string `json:"Image"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		info := Register_struct{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&info)
		Username := info.Username
		Password := info.Password
		Email := info.Email
		if Username == "" || Password == "" || Email == "" {
			fmt.Fprint(w, -2)
		} else {
			if AddUser(Username, Email, Password) {
				if len(info.Image) > 10 {
					img := strings.Split(info.Image, ",")
					file, _ := base64.StdEncoding.DecodeString(img[1])
					User_id := GetUserId(Username)
					os.Remove("Avatar/" + User_id + ".png")
					dst, _ := os.Create("Avatar/" + User_id + ".png")
					dst.Write(file)
					defer dst.Close()
				}
				fmt.Fprint(w, 1)
				fmt.Println("User added")
			} else {
				fmt.Fprint(w, -1)
				fmt.Println("User not added")
			}
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fm, _ := GetFlash(w, r, "message")
		t.ExecuteTemplate(w, "login.html", nil)
		fmt.Fprintf(w, "%s", fm)
	case "POST":
		info := Login_struct{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&info)
		Username := info.Username
		Password := info.Password
		if IsUser(Username, Password) {
			cookie := &http.Cookie{Name: "session", Value: uuid.NewString()}
			AddUUID(Username, cookie.Value)
			http.SetCookie(w, cookie)
			fmt.Fprint(w, 1)
		} else {
			fmt.Fprint(w, -1)
		}
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(r) {
		cookie, _ := r.Cookie("session")
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		SetFlash(w, "message", []byte("You have succesfully logged out"))
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
