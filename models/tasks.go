package models

import (
	"log"

	"github.com/Kamaropoulos/goctapus-mongo/core"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"

	Log "github.com/sirupsen/logrus"
)

// Task is a struct containing Task data
type Task struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items" bson:"items"`
}

// GetTasks from the DB
func GetTasks() TaskCollection {
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("goapp").C("tasks")
	result := TaskCollection{}
	err := c.Find(nil).All(&result.Tasks)

	if err != nil {
		panic(err)
	}

	return result
}

// PutTask into DB
func PutTask(name string) (string, error) {
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("goapp").C("tasks")
	task := Task{ID: bson.NewObjectId(), Name: name}
	err := c.Insert(&task)
	if err != nil {
		log.Fatal(err)
	}
	return task.ID.Hex(), nil
}

// DeleteTask from DB
func DeleteTask(id string) (int64, error) {
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("goapp").C("tasks")
	err := c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		Log.Warn(err)
	}

	return 0, nil
}
