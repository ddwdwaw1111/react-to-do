package main

import (
	"fmt"
	"time"

	"github.com/db/sqldb"
	"github.com/db/taskserver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

func main() {
	db := sqldb.ConnectDB()
	fmt.Print(db)
	//Create a new store and add three test to it.
	tasks := taskserver.NewTaskServer()
	tasks.Store.CreateTask("test0", time.Now().Add(time.Hour*24))
	tasks.Store.CreateTask("test1", time.Now().Add(time.Hour*24))
	tasks.Store.CreateTask("test2", time.Now().Add(time.Hour*24))
	// fmt.Println(alltasks)

	//Start iris serve
	app := iris.New()

	//Router
	tasksAPI := app.Party("/tasks")
	{
		tasksAPI.Use(iris.Compression)

		tasksAPI.Get("/", tasks.GetAll)

		tasksAPI.Post("/", tasks.CreateNew)

		tasksAPI.Patch("/", tasks.UpdateComplete)

		tasksAPI.Delete("/", tasks.DeleteTask)
	}

	app.Listen(":8080")
}
