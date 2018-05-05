package goctapus

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"

	Log "github.com/sirupsen/logrus"
)

// InitDB initializes the connection to the Database
func InitDB(dbString string) *sql.DB {
	db, err := sql.Open("mysql", dbString)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func ConnectDB(host string) *mgo.Session {
	db, err := mgo.Dial(host)
	if err != nil {
		Log.Fatal("CreateSession: %s\n", err)
	}

	db.SetMode(mgo.Monotonic, true)

	return db
}
