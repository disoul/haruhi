package main

import (
	"gopkg.in/mgo.v2"
)

func CreateMongoSession(host string) *mgo.Session {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	return session
}

func saveTask(task Task) error {
	c := MongoSession.DB("haruhi").C("task")
	err := c.Insert(&task)
	if err != nil {
		return err
	}
	return nil
}
