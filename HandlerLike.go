package forum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Like struct {
	Value   string `json:"value"`
	PorM_id string `json:"PorM_id"`
	Table   string `json:"table"`
}

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if CheckCookie(r) {
			user := GetUser(r)
			info := Like{}
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&info)
			LikeOrDisLike := info.Value
			table := info.Table
			PorM_id, _ := strconv.Atoi(info.PorM_id)
			switch table {
			case "LikePost":
				if LikeOrDisLike == "Like" {
					like := strconv.Itoa(LikePost(PorM_id, user.Id))
					send := []string{like, strconv.Itoa(GetDisLikeCountPost(PorM_id))}
					fmt.Fprint(w, send)
				} else {
					dislike := strconv.Itoa(DisLikePost(PorM_id, user.Id))
					send := []string{strconv.Itoa(GetLikeCountPost(PorM_id)), dislike}
					fmt.Fprint(w, send)
				}
			case "LikeComment":
				if LikeOrDisLike == "Like" {
					like := strconv.Itoa(LikeComment(PorM_id, user.Id))
					send := []string{like, strconv.Itoa(GetDisLikeCountComment(PorM_id))}
					fmt.Fprint(w, send)
				} else {
					dislike := strconv.Itoa(DisLikeComment(PorM_id, user.Id))
					send := []string{strconv.Itoa(GetLikeCountComment(PorM_id)), dislike}
					fmt.Fprint(w, send)
				}
			}
		} else {
			SetFlash(w, "message", []byte("You need to be logged in to Like or Dislike"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
