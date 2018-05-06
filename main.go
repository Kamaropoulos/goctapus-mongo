package main

import (
	"os"

	"github.com/Kamaropoulos/goctapus-mongo/core"
	"github.com/Kamaropoulos/goctapus-mongo/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args, "debug", "goapp")

	goctapus.File("/", "public/index.html")
	goctapus.GET("/tasks", handlers.GetTasks())
	goctapus.PUT("/tasks", handlers.PutTask())
	goctapus.DELETE("/tasks/:id", handlers.DeleteTask())

	goctapus.Start()
}
