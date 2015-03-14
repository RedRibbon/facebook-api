package controllers

import (
	"database/sql"
	"facebook-api/app/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"gopkg.in/gorp.v1"
)

func init() {
	revel.OnAppStart(InitDb)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

var InitDb func() = func() {
	if db, err := sql.Open("sqlite3", "/tmp/facebook_db.bin"); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db:      db,
			Dialect: gorp.SqliteDialect{}}
	}
	// Defines the table for use by GORP
	// This is a function we will create soon.
	defineTable(Dbm)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

func defineTable(dbm *gorp.DbMap) {
	dbm.AddTableWithName(models.Feed{}, "feeds").SetKeys(false, "Id")
	dbm.AddTableWithName(models.Comment{}, "comments").SetKeys(false, "Id")
	dbm.AddTableWithName(models.Like{}, "likes").SetKeys(true, "Id")
	dbm.AddTableWithName(models.User{}, "users").SetKeys(false, "Id")
}
