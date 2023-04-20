package forum

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS LikeComment (Comment_id INTEGER, User_id INTEGER, IsLike INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS LikePost (Post_id INTEGER, User_id INTEGER, IsLike INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS ReportPost (Post_id INTEGER UNIQUE, User_id INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS ReportComment (Post_id INTEGER, Comment_id INTEGER UNIQUE, User_id INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS Comment (Id INTEGER UNIQUE PRIMARY KEY, Post_id INTEGER, User_id INTEGER, Text TEXT, CreatedAt DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS Post (Id INTEGER UNIQUE PRIMARY KEY, User_id INTEGER, Title TEXT UNIQUE,Text TEXT, CreatedAt DATETIME, UpdatedAT DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS User (Id INTEGER UNIQUE PRIMARY KEY,UUID TEXT , Username TEXT UNIQUE, Email TEXT UNIQUE, Password TEXT, Role TEXT, CreatedAt DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS Categories (Id INTEGER UNIQUE PRIMARY KEY, Name TEXT)")
	db.Exec("CREATE TABLE IF NOT EXISTS Post_Categories (Post_id INTEGER, Categorie_id INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS RequestMod (User_id INTEGER, Reason TEXT)")
}
