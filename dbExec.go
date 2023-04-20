package forum

import (
	"io"
	"os"
	"time"
)

func AddUser(username string, mail string, password string) bool {
	hashpwd, _ := HashPassword(password)
	CreatedAt := time.Now()
	_, err := db.Exec("INSERT INTO User(UUID,Username,Email,Password,Role,CreatedAt) VALUES (?,?,?,?,?,?)", "NONE", username, mail, hashpwd, "CLASSIC", CreatedAt)
	if err != nil {
		return false
	}
	var User_id string
	db.QueryRow("SELECT Id FROM User WHERE Username = ? AND Email = ?", username, mail).Scan(&User_id)
	fin, _ := os.Open("Avatar/default.png")
	defer fin.Close()
	fout, _ := os.Create("Avatar/" + User_id + ".png")
	io.Copy(fout, fin)
	return true
}

func AddPost(user_id int, Title string, Text string) int {
	CreatedAt := time.Now()
	_, err1 := db.Exec("INSERT INTO Post(User_id,Title,Text,CreatedAt,UpdatedAt) VALUES (?,?,?,?,?)", user_id, Title, Text, CreatedAt, CreatedAt)
	var id int
	db.QueryRow("SELECT Id FROM Post WHERE User_id = ? AND Title = ?", user_id, Title).Scan(&id)
	if err1 != nil {
		return -1
	}
	return id
}

func DeleteRequest(Username string) {
	var User_id int
	db.QueryRow("SELECT Id FROM User WHERE Username = ?", Username).Scan(&User_id)
	db.Exec("DELETE from RequestMod WHERE User_id = ?", User_id)
}

func RankupMod(Username string) {
	db.Exec("UPDATE User Set Role = ? WHERE Username = ?", "MOD", Username)
}
func RankupAdmin(Username string) {
	db.Exec("UPDATE User Set Role = ? WHERE Username = ?", "ADMIN", Username)
}

func AddComment(Post_id int, User_id int, Text string) {
	date := time.Now()
	db.Exec("INSERT INTO Comment(Post_id,User_id,Text,CreatedAt) VALUES (?,?,?,?)", Post_id, User_id, Text, date)
	db.Exec("UPDATE Post SET UpdatedAt = ? WHERE Id = ?", date, Post_id)
}
func AddReportPost(Post_id int, User_id int) {
	db.Exec("INSERT INTO ReportPost(Post_id,User_id) VALUES (?,?)", Post_id, User_id)
}
func AddReportComment(Post_id int, Comment_id int, User_id int) {
	db.Exec("INSERT INTO ReportComment(Post_id,Comment_id,User_id) VALUES (?,?,?)", Post_id, Comment_id, User_id)
}
func EditPost(Post_id int, Title string, Text string) {
	db.Exec("UPDATE Post SET Title = ?, Text = ?, UpdatedAt = ? WHERE Id = ?", Title, Text, time.Now(), Post_id)
}
func DeletePost(Post_id int) {
	db.Exec("DELETE from LikeComment JOIN Comment WHERE LikeComment.Comment_id = Comment.Id AND Comment.Post_id = ?", Post_id)
	db.Exec("DELETE from Comment WHERE Post_id = ?", Post_id)
	db.Exec("DELETE from Post WHERE Id = ?", Post_id)
	db.Exec("DELETE from LikePost WHERE Post_id = ?", Post_id)
	db.Exec("DELETE from ReportPost WHERE Post_id = ?", Post_id)
	db.Exec("DELETE from ReportComment WHERE Post_id = ?", Post_id)
	DeleteCategories(Post_id)
}
func DeleteComment(Comment_id int) {
	db.Exec("DELETE from Comment WHERE Id = ?", Comment_id)
	db.Exec("DELETE from ReportComment WHERE Comment_id = ?", Comment_id)
}
func CancelReportComment(Comment_id int) {
	db.Exec("DELETE from ReportComment WHERE Comment_id = ?", Comment_id)
}
func CancelReportPost(Post_id int) {
	db.Exec("DELETE from ReportPost WHERE Post_id = ?", Post_id)
}
func DisLikePost(Post_id int, User_id int) int {
	IsLike := 0
	db.QueryRow("SELECT IsLike FROM LikePost WHERE Post_id = ? AND User_id = ?", Post_id, User_id).Scan(&IsLike)
	switch IsLike {
	case 0:
		db.Exec("INSERT INTO LikePost VALUES (?,?,?)", Post_id, User_id, -1)
	case 1:
		db.Exec("UPDATE LikePost SET IsLike = ? WHERE Post_id = ? AND User_id = ?", -1, Post_id, User_id)
	case -1:
		db.Exec("DELETE from LikePost WHERE Post_id = ? AND User_id = ?", Post_id, User_id)
	}
	return GetDisLikeCountPost(Post_id)
}
func DisLikeComment(Comment_id int, User_id int) int {
	IsLike := 0
	db.QueryRow("SELECT IsLike FROM LikeComment WHERE Comment_id = ? AND User_id = ?", Comment_id, User_id).Scan(&IsLike)
	switch IsLike {
	case 0:
		db.Exec("INSERT INTO LikeComment VALUES (?,?,?)", Comment_id, User_id, -1)
	case 1:
		db.Exec("UPDATE LikeComment SET IsLike = ? WHERE Comment_id = ? AND User_id = ?", -1, Comment_id, User_id)
	case -1:
		db.Exec("DELETE from LikeComment WHERE Comment_id = ? AND User_id = ?", Comment_id, User_id)
	}
	return GetDisLikeCountComment(Comment_id)
}

func LikePost(Post_id int, User_id int) int {
	IsLike := 0
	db.QueryRow("SELECT IsLike FROM LikePost WHERE Post_id = ? AND User_id = ?", Post_id, User_id).Scan(&IsLike)
	switch IsLike {
	case 0:
		db.Exec("INSERT INTO LikePost VALUES (?,?,?)", Post_id, User_id, 1)
	case -1:
		db.Exec("UPDATE LikePost SET IsLike = ? WHERE Post_id = ? AND User_id = ?", 1, Post_id, User_id)
	case 1:
		db.Exec("DELETE from LikePost WHERE Post_id = ? AND User_id = ?", Post_id, User_id)
	}
	return GetLikeCountPost(Post_id)
}

func LikeComment(Comment_id int, User_id int) int {
	IsLike := 0
	db.QueryRow("SELECT IsLike FROM LikeComment WHERE Comment_id = ? AND User_id = ?", Comment_id, User_id).Scan(&IsLike)
	switch IsLike {
	case 0:
		db.Exec("INSERT INTO LikeComment VALUES (?,?,?)", Comment_id, User_id, 1)
	case -1:
		db.Exec("UPDATE LikeComment SET IsLike = ? WHERE Comment_id = ? AND User_id = ?", 1, Comment_id, User_id)
	case 1:
		db.Exec("DELETE from LikeComment WHERE Comment_id = ? AND User_id = ?", Comment_id, User_id)
	}
	return GetLikeCountComment(Comment_id)
}

func AddRequestMod(User_id int, Reason string) bool {
	_, err := db.Exec("INSERT INTO RequestMod VALUES (?,?)", User_id, Reason)
	return err == nil
}

func AddUUID(Username string, UUID string) {
	db.Exec("UPDATE User SET UUID = ? WHERE Username = ?", UUID, Username)
}

func AddCategories(post_id int, categories []string) {
	for _, categorie := range categories {
		db.Exec("INSERT INTO Post_Categories VALUES (?,?)", post_id, categorie)
	}
}

func DeleteCategories(Post_id int) {
	db.Exec("DELETE from Post_Categories WHERE Post_id = ?", Post_id)
}

func AddCategorie(Title string) {
	db.Exec("INSERT INTO Categories(Name) VALUES (?)", Title)
}
