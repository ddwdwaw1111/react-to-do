package main

import (
	"fmt"
	"time"

	"myapp/backend/internal/taskstore"

	"github.com/kataras/iris/v12"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) getAll(ctx iris.Context) {
	ctx.JSON(ts.store.GetAllTasks())
}

func (ts *taskServer) post(ctx iris.Context) {
	fmt.Println("jinlaile")
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.JSON(task)
}

// func create(ctx iris.Context) {
// 	newTask := Task{"Mastering Concurrency in Go", false}
// 	tasks = append(tasks, newTask)
// 	ctx.JSON(tasks)
// }

func main() {
	tasks := NewTaskServer()
	tasks.store.CreateTask("test0", time.Now().Add(time.Hour*24))
	tasks.store.CreateTask("test1", time.Now().Add(time.Hour*24))
	tasks.store.CreateTask("test2", time.Now().Add(time.Hour*24))
	alltasks := tasks.store.GetAllTasks()
	fmt.Println(alltasks)

	app := iris.New()

	tasksAPI := app.Party("/tasks")
	{
		tasksAPI.Use(iris.Compression)

		tasksAPI.Get("/", tasks.getAll)
		tasksAPI.Post("/", tasks.post)
		// DELETE : http://localhost:8080/books/<taskId>
		// UPDATE : http://localhost:8080/books/<taskId>
	}

	app.Listen(":8080")
}
