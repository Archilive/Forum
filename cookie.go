package forum

import "net/http"

func CheckCookie(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}
	var cpt int
	db.QueryRow("SELECT count(*) FROM User WHERE UUID = ?", cookie.Value).Scan(&cpt)
	return cpt > 0
}
