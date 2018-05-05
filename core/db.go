package goctapus

import (
	"database/sql"
	"time"

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

func ConnectDB(host string, port string, database string, user string, pass string) {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{host + ":" + port},
		Timeout:  60 * time.Second,
		Database: database,
		Username: user,
		Password: pass,
	}

	Database, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		Log.Fatal("CreateSession: %s\n", err)
	}

	Database.SetMode(mgo.Monotonic, true)

	Log.WithField("Database", Database).Debug("Succesfully connected to MongoDB")
}
