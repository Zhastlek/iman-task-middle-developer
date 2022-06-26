package models

type Post struct {
	Id     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type URL struct {
	URL string `json:"url"`
}

type PostsID struct {
	Ids []int32 `json:"ids"`
}
