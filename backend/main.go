package main

import (
	"fmt"
	"time"

	"github.com/db/sqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/internal/taskstore"
	"github.com/kataras/iris/v12"
)

type mypage struct {
	Id int `json:"id"`
}

//server
type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

//handler

func (ts *taskServer) getAll(ctx iris.Context) {
	ctx.JSON(ts.store.GetAllTasks())
}

func (ts *taskServer) createNew(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ts.store.CreateTask(task.Text, time.Now())
	ctx.JSON(ts.store.GetAllTasks())
}

func (ts *taskServer) deleteTask(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ts.store.DeleteTask(task.Id)
	ctx.JSON(ts.store.GetAllTasks())
}

func (ts *taskServer) updateComplete(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ts.store.UpdateTask(task)
	ctx.JSON(ts.store.GetAllTasks())
}

func main() {
	db := sqldb.ConnectDB()
	fmt.Print(db)
	//Create a new store and add three test to it.
	tasks := NewTaskServer(db)
	tasks.store.CreateTask("test0", time.Now().Add(time.Hour*24))
	tasks.store.CreateTask("test1", time.Now().Add(time.Hour*24))
	tasks.store.CreateTask("test2", time.Now().Add(time.Hour*24))
	// fmt.Println(alltasks)

	//Start iris serve
	app := iris.New()

	//Router
	tasksAPI := app.Party("/tasks")
	{
		tasksAPI.Use(iris.Compression)

		tasksAPI.Get("/", tasks.getAll)

		tasksAPI.Post("/", tasks.createNew)

		tasksAPI.Patch("/", tasks.updateComplete)

		tasksAPI.Delete("/", tasks.deleteTask)
		// UPDATE : http://localhost:8080/books/<taskId>
	}

	app.Listen(":8080")
}
