package main

import (
	"gopkg.in/mgo.v2"
)

func CreateMongoCollection(host string) *mgo.Collection {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	c := session.DB("haruhi").C("task")

	return c
}

func saveTask(task Task) error {
	err := MongoTaskCollection.Insert(&task)
	if err != nil {
		return err
	}
	return nil
}
