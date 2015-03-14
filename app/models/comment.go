package models

type Comment struct {
	Id        int64  `db:"id" json:"id"`
	FeedId    string `db:"feed_id" json:"feed_id"`
	From      int64  `db:"from" json:"from"`
	Message   string `db:"message" json:"message"`
	LikeCount int64  `db:"like_count" json:"like_count"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}
