package taskserver

import (
	"database/sql"
	"time"

	"github.com/internal/taskstore"
	"github.com/kataras/iris/v12"
)

type TaskServer struct {
	Store *taskstore.TaskStore
}

func NewTaskServer(db *sql.DB) *TaskServer {
	store := taskstore.New(db)
	return &TaskServer{Store: store}
}

//handler

func (ts *TaskServer) GetAll(ctx iris.Context) {
	ctx.JSON(ts.Store.GetAllTasks())

}

func (ts *TaskServer) CreateNew(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ts.Store.CreateTask(task.Text, time.Now())
	ctx.JSON(ts.Store.GetAllTasks())
}

func (ts *TaskServer) DeleteTask(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	// ts.Store.DeleteTask(task.Id)
	// ctx.JSON(ts.Store.GetAllTasks())
}

func (ts *TaskServer) UpdateComplete(ctx iris.Context) {
	var task taskstore.Task
	err := ctx.ReadJSON(&task)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	// ts.Store.UpdateTask(task)
	// ctx.JSON(ts.Store.GetAllTasks())
}
