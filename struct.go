package forum

import "time"

type TemplateSetting struct {
	IsLogged bool
	User     User
}
type TemplateEdit struct {
	IsLogged   bool
	User       User
	Post       Post
	Categories []Categorie
}
type TemplateCreatePost struct {
	IsLogged   bool
	User       User
	Categories []Categorie
}
type TemplateProfile struct {
	IsLogged bool
	User     User
	UserInfo UserInfo
}

type TemplatePost struct {
	IsLogged   bool
	User       User
	Filter     string
	Order      string
	Categorie  string
	Posts      []Post
	Categories []Categorie
}
type TemplateComment struct {
	IsLogged bool
	User     User
	Post     Post
	Comments []Comment
}
type TemplateAdmin struct {
	IsLogged   bool
	User       User
	ReportInfo ReportInfo
	RequestMod []Request
}
type Request struct {
	User_id  int
	Username string
	Reason   string
}
type Post struct {
	Id           int
	User_id      int
	Title        string
	Text         string
	NbrComments  int
	CreatedAt    time.Time
	StrCreatedAt string
	UpdatedAt    time.Time
	StrUpdatedAt string
	LikeCount    int
	DisLikeCount int
	IsLiked      bool
	IsDisLiked   bool
	Username     string
	Categories   []Categorie
}

type Comment struct {
	Id           int
	Post_id      int
	User_id      int
	Text         string
	CreatedAt    time.Time
	StrCreatedAt string
	LikeCount    int
	DisLikeCount int
	IsLiked      bool
	IsDisLiked   bool
	Username     string
}
type User struct {
	Id           int
	UUID         string
	Username     string
	Email        string
	Password     string
	CreatedAt    time.Time
	StrCreatedAt string
	Role         string
	IsMod        bool
	IsAdmin      bool
}
type UserInfo struct {
	User          User
	PostedPost    []Post
	PostedComment []Comment
	LikedPost     []Post
	NbrLike       int
	NbrDisLike    int
}
type Categorie struct {
	Id   int
	Name string
}
type ReportInfo struct {
	ReportPosts    []Post
	ReportComments []Comment
}
