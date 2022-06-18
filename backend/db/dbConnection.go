package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type tasks struct {
	Id     int
	Text   string
	DuoDay time.Time
}

func main() {
	fmt.Print("database connecting")
	db, err := sql.Open("mysql", "mailtrain:mailtrain@tcp(zihao@4.7.168.253:22)/tasks")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT id, text FROM tasks")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tasks tasks
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tasks.Id, &tasks.Text)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tasks.Text)
	}
}
