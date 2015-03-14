package controllers

import (
	"facebook-api/app/models"
	"strconv"
	"time"

	"github.com/revel/revel"
)

type ApiCtrl struct {
	GorpController
}

type ErrResponse struct {
	Message string `json:"message"`
}

func (c ApiCtrl) MostCommentedFeeds() revel.Result {
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(20))

	from := c.Params.Get("from")
	if len(from) > 0 {
		if !checkDateStr(from) {
			revel.INFO.Printf("from '%v' is wrong format", from)
			c.Response.Status = 400
			result := ErrResponse{
				Message: "from is wrong format"}
			return c.RenderJson(result)
		}
		from = convToSqlDateStr(from, true)
	}

	to := c.Params.Get("to")
	if len(to) > 0 {
		if !checkDateStr(to) {
			revel.INFO.Printf("to '%v' is wrong format", to)
			c.Response.Status = 400
			result := ErrResponse{
				Message: "to is wrong format"}
			return c.RenderJson(result)
		}
		to = convToSqlDateStr(to, false)
	}

	sql := "select A.comment_count CommentCount, feeds.id Id, feeds.[from] [From], feeds.message Message, created_at CreatedAt, updated_at UpdatedAt from feeds inner join (select count(*) as comment_count, feed_id from comments group by feed_id) as A on feeds.id = A.feed_id"
	where := ""

	if len(from) > 0 {
		where += " and feeds.created_at >= strftime('%s', :from)"
	}

	if len(to) > 0 {
		where += " and feeds.created_at < strftime('%s', :to)"
	}

	if len(where) > 0 {
		sql += " where " + where[5:]
	}
	sql = "select users.name FromName, B.* from users inner join (" + sql + ") as B on users.id = B.[From]"
	sql += " order by CommentCount desc limit :limit"

	revel.INFO.Println(sql)
	revel.INFO.Printf("%v %v %v", from, to, limit)

	var feeds []models.FeedCommentView
	_, err := c.Txn.Select(&feeds, sql, map[string]interface{}{
		"from":  from,
		"to":    to,
		"limit": limit,
	})

	if err != nil {
		revel.INFO.Println(err)
		result := ErrResponse{
			Message: "Error trying to get records from DB."}

		return c.RenderJson(result)
	}

	return c.RenderJson(feeds)
}

// YYYY, YYYYMM, YYYYMMDD
func checkDateStr(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}

	length := len(str)
	if length != 4 && length != 6 && length != 8 {
		return false
	}

	switch len(str) {
	case 4:
		str += "0101"
	case 6:
		str += "01"
	}

	if _, err := time.Parse("20060102", str); err != nil {
		return false
	}

	return true
}

func convToSqlDateStr(str string, start bool) string {
	length := len(str)

	switch length {
	case 4:
		str += "0101"
	case 6:
		str += "01"
	}

	t, _ := time.Parse("20060102", str)

	if !start {
		// foudn out end date
		switch length {
		case 4:
			t = t.AddDate(1, 0, 0)
		case 6:
			t = t.AddDate(0, 1, 0)
		}
	}

	return t.Format("2006-01-02")
}
