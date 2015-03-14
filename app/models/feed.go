package models

type Feed struct {
	Id        string `db:"id" json:"id"`
	From      int64  `db:"from" json:"from"`
	Message   string `db:"message" json:"message"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

type FeedCommentView struct {
	Id           string `json:"id"`
	From         int64  `json:"from"`
	FromName     string `json:"from_name"`
	Message      string `json:"message"`
	CommentCount int64  `json:"comment_count"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
