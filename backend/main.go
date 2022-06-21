package main

import (
	"fmt"

	"github.com/db/sqldb"
	"github.com/db/taskserver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

func main() {
	db := sqldb.ConnectDB()
	fmt.Print(db)
	//Create a new store and add three test to it.
	tasks := taskserver.NewTaskServer(db)

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
