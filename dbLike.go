package forum

import (
	"database/sql"
)

func GetLikeCountPost(Post_id int) int {
	var count int
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikePost WHERE IsLike = ? AND Post_id = ?", 1, Post_id).Scan(&count)
	defer db.Close()
	return count
}

func GetDisLikeCountPost(Post_id int) int {
	db, _ := sql.Open("sqlite3", "data.db")
	var count int
	db.QueryRow("SELECT count(*) FROM LikePost WHERE IsLike = ? AND Post_id = ?", -1, Post_id).Scan(&count)
	defer db.Close()
	return count
}

func GetLikeCountComment(Comment_id int) int {
	var count int
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikeComment WHERE IsLike = ? AND Comment_id = ?", 1, Comment_id).Scan(&count)
	defer db.Close()
	return count
}

func GetDisLikeCountComment(Comment_id int) int {
	db, _ := sql.Open("sqlite3", "data.db")
	var count int
	db.QueryRow("SELECT count(*) FROM LikeComment WHERE IsLike = ? AND Comment_id = ?", -1, Comment_id).Scan(&count)
	defer db.Close()
	return count
}

func IsLikedPost(Post_id int, user_id int) bool {
	Like := 0
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikePost WHERE Post_id = ? AND User_id = ? AND IsLike = ?", Post_id, user_id, 1).Scan(&Like)
	defer db.Close()
	return Like == 1
}

func IsDisLikedPost(Post_id int, user_id int) bool {
	Like := 0
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikePost WHERE Post_id = ? AND User_id = ? AND IsLike = ?", Post_id, user_id, -1).Scan(&Like)
	defer db.Close()
	return Like == 1
}
func IsLikedComment(Comment_id int, user_id int) bool {
	Like := 0
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikeComment WHERE Comment_id = ? AND User_id = ? AND IsLike = ?", Comment_id, user_id, 1).Scan(&Like)
	defer db.Close()
	return Like == 1
}

func IsDisLikedComment(Comment_id int, user_id int) bool {
	Like := 0
	db, _ := sql.Open("sqlite3", "data.db")
	db.QueryRow("SELECT count(*) FROM LikeComment WHERE Comment_id = ? AND User_id = ? AND IsLike = ?", Comment_id, user_id, -1).Scan(&Like)
	defer db.Close()
	return Like == 1
}
