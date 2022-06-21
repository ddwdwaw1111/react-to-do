package taskstore

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Task struct {
	Id         int       `json:"id"`
	Text       string    `json:"text"`
	Due        time.Time `json:"due"`
	IsComplete bool      `json:"isComplete"`
}

// TaskStore is a simple in-memory database of tasks; TaskStore methods are
// safe to call concurrently.
type TaskStore struct {
	db *sql.DB
}

func New(db *sql.DB) *TaskStore {
	fmt.Println("Database Stored")
	return &TaskStore{db: db}
}

func (ts *TaskStore) GetAllTasks() []Task {
	var tasks []Task
	res, err := ts.db.Query("SELECT task_id,text FROM tasks")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var task Task
		err := res.Scan(&task.Id, &task.Text)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}
	return tasks
}

// CreateTask creates a new task in the store.
func (ts *TaskStore) CreateTask(text string, due time.Time) {

	// task.Tags = make([]string, len(tags))
	// copy(task.Tags, tags)
	sql := "INSERT INTO tasks(text,duo_date) VALUE (?,?)"
	res, err := ts.db.Exec(sql, text, due)

	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}

// GetTask retrieves a task from the store, by id. If no such id exists, an
// error is returned.
// func (ts *TaskStore) GetTask(id int) (Task, error) {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	t, ok := ts.tasks[id]
// 	if ok {
// 		return t, nil
// 	} else {
// 		return Task{}, fmt.Errorf("task with id=%d not found", id)
// 	}
// }

// func (ts *TaskStore) UpdateTask(newTask Task) (Task, error) {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	t, ok := ts.tasks[newTask.Id]
// 	if ok {
// 		ts.tasks[newTask.Id] = newTask
// 		fmt.Println(ts)
// 		return t, nil
// 	} else {
// 		return Task{}, fmt.Errorf("task with id=%d not found", newTask.Id)
// 	}
// }

// // DeleteTask deletes the task with the given id. If no such id exists, an error
// // is returned.
// func (ts *TaskStore) DeleteTask(id int) error {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	if _, ok := ts.tasks[id]; !ok {
// 		return fmt.Errorf("task with id=%d not found", id)
// 	}

// 	delete(ts.tasks, id)
// 	return nil
// }

// // DeleteAllTasks deletes all tasks in the store.
// func (ts *TaskStore) DeleteAllTasks() error {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	ts.tasks = make(map[int]Task)
// 	return nil
// }

// GetAllTasks returns all the tasks in the store, in arbitrary order.

// GetTasksByTag returns all the tasks that have the given tag, in arbitrary
// order.
// func (ts *TaskStore) GetTasksByTag(tag string) []Task {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	var tasks []Task

// taskloop:
// 	for _, task := range ts.tasks {
// 		for _, taskTag := range task.Tags {
// 			if taskTag == tag {
// 				tasks = append(tasks, task)
// 				continue taskloop
// 			}
// 		}
// 	}
// 	return tasks
// }

// GetTasksByDueDate returns all the tasks that have the given due date, in
// arbitrary order.
// func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	var tasks []Task

// 	for _, task := range ts.tasks {
// 		y, m, d := task.Due.Date()
// 		if y == year && m == month && d == day {
// 			tasks = append(tasks, task)
// 		}
// 	}

// 	return tasks
// }
