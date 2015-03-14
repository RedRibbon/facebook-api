package models

type Like struct {
	Id     int64  `db:"id" json:"id"`
	UserId int64  `db:"user_id" json:"user_id"`
	FeedId string `db:"feed_id" json:"feed_id"`
}
