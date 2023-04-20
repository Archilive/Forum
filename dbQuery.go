package forum

import (
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetNbrComment(post_id int) int {
	var count int
	db.QueryRow("SELECT count(*) FROM Comment WHERE Post_id = ?", post_id).Scan(&count)
	return count
}

func GetAllPost(user_id int) []Post {
	rows, _ := db.Query("SELECT * FROM Post")
	var AllPost []Post
	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Text, &post.CreatedAt, &post.UpdatedAt)
		post.StrCreatedAt = TimetoString(post.CreatedAt)
		post.StrUpdatedAt = TimetoString(post.UpdatedAt)
		post.LikeCount = GetLikeCountPost(post.Id)
		post.DisLikeCount = GetDisLikeCountPost(post.Id)
		post.NbrComments = GetNbrComment(post.Id)
		post.Categories = GetCategories(post.Id)
		db.QueryRow("SELECT Username FROM User WHERE Id = ?", post.User_id).Scan(&post.Username)
		if user_id >= 0 {
			post.IsLiked = IsLikedPost(post.Id, user_id)
			post.IsDisLiked = IsDisLikedPost(post.Id, user_id)
		}
		AllPost = append(AllPost, post)
	}
	return AllPost
}

func GetPost(id int, user_id int) Post {
	var post Post
	db.QueryRow("SELECT * FROM Post WHERE Id = ?", id).Scan(&post.Id, &post.User_id, &post.Title, &post.Text, &post.CreatedAt, &post.UpdatedAt)
	post.LikeCount = GetLikeCountPost(id)
	post.StrCreatedAt = TimetoString(post.CreatedAt)
	post.DisLikeCount = GetDisLikeCountPost(id)
	db.QueryRow("SELECT Username FROM User WHERE Id = ?", post.User_id).Scan(&post.Username)
	post.Categories = GetCategories(post.Id)
	if user_id >= 0 {
		post.IsLiked = IsLikedPost(post.Id, user_id)
		post.IsDisLiked = IsDisLikedPost(post.Id, user_id)
	}
	return post
}

func GetAllComment(Post_id int, user_id int) []Comment {
	rows, _ := db.Query("SELECT * FROM Comment WHERE Post_id = ?", Post_id)
	var AllComment []Comment
	for rows.Next() {
		var Comment Comment
		rows.Scan(&Comment.Id, &Comment.Post_id, &Comment.User_id, &Comment.Text, &Comment.CreatedAt)
		Comment.LikeCount = GetLikeCountComment(Comment.Id)
		Comment.DisLikeCount = GetDisLikeCountComment(Comment.Id)
		Comment.StrCreatedAt = TimetoString(Comment.CreatedAt)
		db.QueryRow("SELECT Username FROM User WHERE Id = ?", Comment.User_id).Scan(&Comment.Username)
		if user_id >= 0 {
			Comment.IsLiked = IsLikedComment(Comment.Id, user_id)
			Comment.IsDisLiked = IsDisLikedComment(Comment.Id, user_id)
		}
		AllComment = append(AllComment, Comment)
	}
	return AllComment
}

func IsUser(Username string, password string) bool {
	var passworduser string
	db.QueryRow("SELECT Password FROM User WHERE Username = ?", Username).Scan(&passworduser)
	return CheckPasswordHash(password, passworduser)
}

func IsOwner(Post_id int, User_id int) bool {
	var nbr int
	db.QueryRow("SELECT EXicount(*) FROM Post WHERE User_id = ? AND Id = ?", User_id, Post_id).Scan(&nbr)
	return nbr > 0
}
func GetPostId(Comment_id int) int {
	var Post_id int
	db.QueryRow("SELECT Post.Id FROM Post JOIN Comment WHERE Post.Id = Comment.Post_id AND Comment.Id = ?", Comment_id).Scan(&Post_id)
	return Post_id
}

func GetRequest() []Request {
	var Requests []Request
	rows, _ := db.Query("SELECT User.Username,RequestMod.Reason FROM User JOIN RequestMod WHERE User.Id = RequestMod.User_id")
	for rows.Next() {
		var request Request
		rows.Scan(&request.Username, &request.Reason)
		Requests = append(Requests, request)
	}
	return Requests
}

func GetUser(r *http.Request) User {
	var user User
	cookie, _ := r.Cookie("session")
	db.QueryRow("SELECT * FROM User WHERE UUID = ?", cookie.Value).
		Scan(&user.Id, &user.UUID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	switch user.Role {
	case "MOD":
		user.IsMod, user.IsAdmin = true, false
	case "ADMIN":
		user.IsMod, user.IsAdmin = true, true
	case "CLASSIC":
		user.IsMod, user.IsAdmin = false, false
	}
	return user
}
func GetUserId(username string) string {
	var id int
	db.QueryRow("SELECT Id FROM User WHERE Username = ?", username).Scan(&id)
	return strconv.Itoa(id)
}

func GetAllUserInfo(User_id int) UserInfo {
	var userinfo UserInfo
	var user User
	db.QueryRow("SELECT * FROM User WHERE Id = ?", User_id).
		Scan(&user.Id, &user.UUID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	user.StrCreatedAt = TimetoString(user.CreatedAt)
	userinfo.User = user
	rows, _ := db.Query("SELECT * FROM Post WHERE User_id = ?", userinfo.User.Id)
	for rows.Next() {
		var Post Post
		rows.Scan(&Post.Id, &Post.User_id, &Post.Title, &Post.Text, &Post.CreatedAt, &Post.UpdatedAt)
		Post.StrUpdatedAt = TimetoString(Post.UpdatedAt)
		Post.StrCreatedAt = TimetoString(Post.CreatedAt)
		Post.LikeCount = GetLikeCountPost(Post.Id)
		Post.DisLikeCount = GetDisLikeCountPost(Post.Id)
		Post.NbrComments = GetNbrComment(Post.Id)
		userinfo.PostedPost = append(userinfo.PostedPost, Post)

	}
	rows, _ = db.Query("SELECT Post.* FROM Post Join LikePost WHERE Post.Id = LikePost.Post_id AND LikePost.IsLike = 1 AND LikePost.User_id = ?", User_id)
	for rows.Next() {
		var Post Post
		rows.Scan(&Post.Id, &Post.User_id, &Post.Title, &Post.Text, &Post.CreatedAt, &Post.UpdatedAt)
		Post.StrUpdatedAt = TimetoString(Post.UpdatedAt)
		Post.StrCreatedAt = TimetoString(Post.CreatedAt)
		Post.LikeCount = GetLikeCountPost(Post.Id)
		Post.DisLikeCount = GetDisLikeCountPost(Post.Id)
		Post.NbrComments = GetNbrComment(Post.Id)
		userinfo.LikedPost = append(userinfo.LikedPost, Post)
	}
	rows, _ = db.Query("SELECT * FROM Comment WHERE User_id = ?", userinfo.User.Id)
	for rows.Next() {
		var Comment Comment
		rows.Scan(&Comment.Id, &Comment.Post_id, &Comment.User_id, &Comment.Text, &Comment.CreatedAt)
		Comment.StrCreatedAt = TimetoString(Comment.CreatedAt)
		Comment.LikeCount = GetLikeCountComment(Comment.Id)
		Comment.DisLikeCount = GetDisLikeCountComment(Comment.Id)
		userinfo.PostedComment = append(userinfo.PostedComment, Comment)
	}
	cpt := 0
	db.QueryRow("SELECT count(*) FROM Comment JOIN LikeComment WHERE Comment.User_id = ? AND Comment.Id = LikeComment.Comment_id AND LikeComment.IsLike = 1", userinfo.User.Id).Scan(&cpt)
	userinfo.NbrLike += cpt
	cpt = 0
	db.QueryRow("SELECT count(*) FROM Post JOIN LikePost WHERE Post.User_id = ? AND Post.Id = LikePost.Post_id AND LikePost.IsLike = 1", userinfo.User.Id).Scan(&cpt)
	userinfo.NbrLike += cpt
	db.QueryRow("SELECT count(*) FROM Comment JOIN LikeComment WHERE Comment.User_id = ? AND Comment.Id = LikeComment.Comment_id AND LikeComment.IsLike = -1", userinfo.User.Id).Scan(&cpt)
	userinfo.NbrDisLike += cpt
	cpt = 0
	db.QueryRow("SELECT count(*) FROM Post JOIN LikePost WHERE Post.User_id = ? AND Post.Id = LikePost.Post_id AND LikePost.IsLike = -1", userinfo.User.Id).Scan(&cpt)
	userinfo.NbrDisLike += cpt
	return userinfo
}

func GetReport() ReportInfo {
	var ReportInfo ReportInfo
	rows, _ := db.Query("SELECT Post.*,User.Username FROM Post JOIN ReportPost,User ON Post.Id = ReportPost.Post_id AND Post.User_id = User.Id ")
	for rows.Next() {
		var Post Post
		rows.Scan(&Post.Id, &Post.User_id, &Post.Title, &Post.Text, &Post.CreatedAt, &Post.UpdatedAt, &Post.Username)
		Post.StrCreatedAt = TimetoString(Post.CreatedAt)
		Post.StrUpdatedAt = TimetoString(Post.UpdatedAt)
		Post.NbrComments = GetNbrComment(Post.Id)
		ReportInfo.ReportPosts = append(ReportInfo.ReportPosts, Post)
	}
	rows, _ = db.Query("SELECT Comment.*,User.Username FROM Comment JOIN ReportComment,User ON Comment.Id = ReportComment.Comment_id AND Comment.User_id = User.Id ")
	for rows.Next() {
		var Comment Comment
		rows.Scan(&Comment.Id, &Comment.Post_id, &Comment.User_id, &Comment.Text, &Comment.CreatedAt, &Comment.Username)
		Comment.StrCreatedAt = TimetoString(Comment.CreatedAt)
		ReportInfo.ReportComments = append(ReportInfo.ReportComments, Comment)
	}
	return ReportInfo
}

func GetALLCategories() []Categorie {
	var Categories []Categorie
	rows, _ := db.Query("SELECT * FROM Categories")
	for rows.Next() {
		var categorie Categorie
		rows.Scan(&categorie.Id, &categorie.Name)
		Categories = append(Categories, categorie)
	}
	return Categories
}

func GetCategories(Post_id int) []Categorie {
	var Categories []Categorie
	rows, _ := db.Query("SELECT Categories.* FROM Categories JOIN Post_Categories WHERE Post_Categories.Categorie_id = Categories.Id AND Post_Categories.Post_id = ? ", Post_id)
	for rows.Next() {
		var categorie Categorie
		rows.Scan(&categorie.Id, &categorie.Name)
		Categories = append(Categories, categorie)
	}
	return Categories
}
