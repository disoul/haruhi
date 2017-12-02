package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresConfig struct {
	host     string
	user     string
	dbname   string
	sslmode  string
	password string
}

func getDbConnection(config PostgresConfig) (*gorm.DB, error) {
	configString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		config.host, config.user, config.dbname,
		config.sslmode, config.password,
	)

	db, err := gorm.Open("postgres", configString)
	if err != nil {
		return db, err
	}

	dbInit(db)
	return db, nil
}

func dbInit(db *gorm.DB) {
	db.AutoMigrate(&Task{})
	if (!db.HasTable(&Task{})) {
		db.CreateTable(&Task{})
	}
}

func saveTask(querytask Task) (*Task, error) {
	//var task Task
	//err := PostgresDb.Where(Task{Name: querytask.Name}).Assign(querytask).FirstOrCreate(&task).Error
	fmt.Printf("\ndb %v", PostgresDb)
	if PostgresDb.NewRecord(querytask) {
		fmt.Print("no record")
		PostgresDb.Create(&querytask)
	}

	return &querytask, nil
}
