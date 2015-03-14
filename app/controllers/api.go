package controllers

import (
	"errors"
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
	to := c.Params.Get("to")

	var err error

	if from, err = sanitizeDateStr(from, true); err != nil {
		return c.makeErrResponse("from is worng format")
	}

	if to, err = sanitizeDateStr(to, false); err != nil {
		return c.makeErrResponse("to is worng format")
	}

	var sql = "select A.comment_count CommentCount, feeds.id Id, feeds.[from] [From], feeds.message Message, created_at CreatedAt, updated_at UpdatedAt from feeds inner join (select count(*) as comment_count, feed_id from comments group by feed_id) as A on feeds.id = A.feed_id"
	var where string

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
	_, err = c.Txn.Select(&feeds, sql, map[string]interface{}{
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

func (c ApiCtrl) MostPostedPersons() revel.Result {
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(20))
	from := c.Params.Get("from")
	to := c.Params.Get("to")

	var err error

	if from, err = sanitizeDateStr(from, true); err != nil {
		return c.makeErrResponse("from is worng format")
	}

	if to, err = sanitizeDateStr(to, false); err != nil {
		return c.makeErrResponse("to is worng format")
	}

	var sql = "select [from], count(*) count from feeds"
	var where string

	if len(from) > 0 {
		where += " and created_at >= strftime('%s', :from)"
	}

	if len(to) > 0 {
		where += " and created_at < strftime('%s', :to)"
	}

	if len(where) > 0 {
		sql += " where " + where[5:]
	}

	sql += " group by [from]"
	sql = "select users.id Id, users.name Name, A.count Count from users inner join (" + sql + " ) as A on users.id = A.[from]"
	sql += " order by Count desc limit :limit"

	revel.INFO.Println(sql)
	revel.INFO.Printf("%v %v %v", from, to, limit)

	var feeds []models.UserPostView
	_, err = c.Txn.Select(&feeds, sql, map[string]interface{}{
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

func (c ApiCtrl) makeErrResponse(msg string) revel.Result {
	c.Response.Status = 400
	result := ErrResponse{Message: msg}
	return c.RenderJson(result)
}

func sanitizeDateStr(str string, start bool) (string, error) {
	if len(str) > 0 {
		if !checkDateStr(str) {
			return "", errors.New("wrong date format")
		}
		str = convToSqlDateStr(str, start)
	}

	return str, nil
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
