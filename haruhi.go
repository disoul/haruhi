package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var PostgresDb *gorm.DB

func main() {
	var err error
	PostgresDb, err = getDbConnection(
		PostgresConfig{
			host:     "127.0.0.1",
			user:     "postgres",
			password: "postgres",
			sslmode:  "disable",
			dbname:   "haruhi",
		},
	)
	fmt.Printf("%v", PostgresDb)
	if err != nil {
		fmt.Errorf("connect to db error")
		panic(err)
	}
	defer PostgresDb.Close()

	initTask()

	CreateManagerServer(7777)
}

func initTask() {
	RegisteredTasks = make(map[string]*Task)
	HaruhiTaskQuery = make(map[string]*TaskQuery)
}
